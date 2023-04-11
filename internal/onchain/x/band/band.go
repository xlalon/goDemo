package band

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/internal/onchain/conf"
	"github.com/xlalon/golee/pkg/json"
	"github.com/xlalon/golee/pkg/math/decimal"
	"github.com/xlalon/golee/pkg/math/rand"
	"github.com/xlalon/golee/pkg/math/sort"
	"github.com/xlalon/golee/pkg/net/http/client"
)

type Band struct {
	*onchain.Chain

	nodeClient client.Client

	maxScanHeightRange int64
}

func New(conf *conf.ChainConfig) *Band {
	return &Band{
		Chain: &onchain.Chain{
			Code:   "BAND",
			Config: conf,
		},

		nodeClient:         client.NewRestyClient(conf.NodeUrl),
		maxScanHeightRange: 2,
	}
}

func (b *Band) GetLatestHeight(ctx context.Context) (int64, error) {
	resp, err := b.nodeClient.Get(ctx, "/blocks/latest", nil)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	height := json.JGet(resp, "block.header.height")
	return height.Int(), nil
}

func (b *Band) GetTxnByHash(ctx context.Context, txHash string) ([]*onchain.Transaction, error) {
	var txs []*onchain.Transaction

	resp, err := b.nodeClient.Get(ctx, fmt.Sprintf("/cosmos/tx/v1beta1/txs/%s", txHash), nil)
	if err != nil {
		fmt.Println(err)
		return txs, err
	}

	txs, _ = b.parseTxResponse(json.JGet(resp, "tx_response").String())

	err = b.updateTxsStatus(ctx, txs)

	return txs, err
}

func (b *Band) ScanTxn(ctx context.Context, cursor *onchain.Cursor) ([]*onchain.Transaction, error) {
	latestHeight, err := b.GetLatestHeight(ctx)
	if err != nil {
		return nil, err
	}
	var txs []*onchain.Transaction
	if cursor.Height == 0 {
		cursor.Height = latestHeight
		return txs, nil
	}
	maxHeight := sort.Min([]int64{latestHeight, cursor.Height + b.maxScanHeightRange})
	for height := cursor.Height + 1; height < maxHeight+1; height++ {
		_txs, errScan := b.scanTxnByBlock(ctx, height)
		if errScan != nil {
			return nil, errScan
		}
		txs = append(txs, _txs...)
	}

	cursor.Height = maxHeight

	return txs, nil
}

func (b *Band) scanTxnByBlock(ctx context.Context, height int64) ([]*onchain.Transaction, error) {

	var txs []*onchain.Transaction

	params := map[string]string{
		"events":            fmt.Sprintf("tx.height=%d", height),
		"pagination.offset": strconv.Itoa(1),
		"pagination.limit":  strconv.Itoa(100),
	}

	resp, err := b.nodeClient.Get(ctx, "/cosmos/tx/v1beta1/txs", params)
	if err != nil {
		return txs, err
	}
	txsResp := json.JGet(resp, "tx_responses")

	for _, txResp := range txsResp.Array() {
		_txs, _ := b.parseTxResponse(txResp.String())
		txs = append(txs, _txs...)
	}
	return txs, nil
}

func (b *Band) NewAccount(ctx context.Context, label onchain.Label) (*onchain.Account, error) {
	_ = ctx
	account := &onchain.Account{}
	if label == onchain.AccountDeposit {
		account = &onchain.Account{
			Chain:   b.Code,
			Address: b.Config.DepositAddress,
			Label:   onchain.AccountDeposit,
			Memo:    rand.DigitalMemo(),
		}
	} else if label == onchain.AccountHot {
		account = &onchain.Account{
			Chain:   b.Code,
			Address: b.Config.HotAddress,
			Label:   onchain.AccountHot,
		}
	}
	return account, nil
}

func (b *Band) GetAccount(ctx context.Context, address string) (*onchain.Account, error) {
	// sequence info
	url := fmt.Sprintf("/cosmos/auth/v1beta1/accounts/%s", address)
	resp, err := b.nodeClient.Get(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	sequence := json.JGet(resp, "account.sequence").Int()
	account := &onchain.Account{
		Chain:    b.Code,
		Address:  address,
		Label:    onchain.AccountUnknown,
		Sequence: sequence,
	}

	return account, nil
}

func (b *Band) GetBalance(ctx context.Context, account *onchain.Account, identity string) (decimal.Decimal, error) {
	zero := decimal.Zero()
	url := fmt.Sprintf("/cosmos/bank/v1beta1/balances/%s", account.Address)
	resp, err := b.nodeClient.Get(ctx, url, nil)
	if err != nil {
		return zero, err
	}
	for _, balances := range json.JGet(resp, "balances").Array() {
		if json.JGet(balances.String(), "denom").String() == identity {
			return decimal.NewFromString(json.JGet(balances.String(), "amount").String())
		}
	}
	return zero, errors.New("identity not found")
}

func (b *Band) EstimateFee(ctx context.Context, reqData *onchain.TransferCommand) (*onchain.Fee, error) {
	_ = ctx
	_ = reqData
	return &onchain.Fee{}, nil
}

func (b *Band) Transfer(ctx context.Context, reqData *onchain.TransferCommand) (*onchain.Receipt, error) {
	_ = ctx
	_ = reqData
	return &onchain.Receipt{}, nil
}

func (b *Band) parseTxResponse(txResponse string) ([]*onchain.Transaction, error) {

	var txs []*onchain.Transaction

	status := onchain.TxnStatus{Result: onchain.TxnPending}
	if code := json.JGet(txResponse, "code").Int(); code != 0 {
		status.Result = onchain.TxnFailed
	}

	tx := json.JGet(txResponse, "tx").String()
	if txType := json.JGet(tx, "@type").String(); txType != "/cosmos.tx.v1beta1.Tx" {
		return txs, errors.New("tx type not /cosmos.tx.v1beta1.Tx")
	}
	txHash := json.JGet(txResponse, "txhash").String()
	height := json.JGet(txResponse, "height").Int()
	body := json.JGet(tx, "body").String()
	memo := json.JGet(body, "memo").String()
	for _, messageRaw := range json.JGet(body, "messages").Array() {
		message := messageRaw.String()
		sender := json.JGet(message, "from_address").String()
		receiver := json.JGet(message, "to_address").String()
		if receiver == "" {
			continue
		}
		amountInfo := json.JGet(message, "amount")
		for _, amountRaw := range amountInfo.Array() {
			_amountInfo := amountRaw.String()
			identity := json.JGet(_amountInfo, "denom").String()
			amountInt, err := strconv.ParseInt(json.JGet(_amountInfo, "amount").String(), 10, 64)
			if err != nil || identity == "" || amountInt <= 0 {
				continue
			}
			txn := &onchain.Transaction{
				TxnId: onchain.TxnId{
					Chain:  b.Code,
					TxHash: txHash,
					VOut:   0,
				},
				Receiver: onchain.Receiver{
					Address: receiver,
					Memo:    memo,
				},
				CoinValue: onchain.CoinValue{
					Identity: identity,
					Amount:   decimal.NewFromInt(amountInt),
				},
				Status: status,
				Sender: sender,
				Height: height,
			}
			txs = append(txs, txn)
		}
	}
	return txs, nil
}

func (b *Band) updateTxsStatus(ctx context.Context, txs []*onchain.Transaction) error {
	latestHeight, _ := b.GetLatestHeight(ctx)
	for _, txn := range txs {
		if txn.Status.Result == onchain.TxnFailed {
			continue
		}
		if txn.Height == 0 {
			// update txn height first
			continue
		}
		confirm := latestHeight - txn.Height
		if confirm <= 0 {
			continue
		}
		txn.Status.Confirmations = confirm
		if confirm >= b.Config.IrreversibleBlock {
			txn.Status.Result = onchain.TxnSuccess
		}
	}
	return nil
}
