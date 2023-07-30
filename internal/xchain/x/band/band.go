package band

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/xlalon/golee/internal/xchain"
	"github.com/xlalon/golee/internal/xchain/conf"
	"github.com/xlalon/golee/pkg/json"
	"github.com/xlalon/golee/pkg/math/decimal"
	"github.com/xlalon/golee/pkg/math/rand"
	"github.com/xlalon/golee/pkg/math/sort"
	"github.com/xlalon/golee/pkg/net/http/client"
)

type Band struct {
	*xchain.X

	nodeClient         client.Client
	maxScanHeightRange int64
}

func New(conf *conf.ChainConfig) *Band {
	return &Band{
		X: &xchain.X{
			Code:   "BAND",
			Config: conf,
		},

		nodeClient:         client.NewRestyClient(conf.NodeUrl),
		maxScanHeightRange: 2,
	}
}

func (b *Band) GetLatestHeight(ctx context.Context) (int64, error) {
	resp, err := b.getLatestHeight(ctx)
	if err != nil {
		return -1, err
	}
	height := json.JGet(resp, "block.header.height")
	return height.Int(), nil
}

func (b *Band) getLatestHeight(ctx context.Context) (string, error) {
	return b.nodeClient.Get(ctx, "/blocks/latest", nil)
}

func (b *Band) GetTransfersByHash(ctx context.Context, txHash string) ([]*xchain.Transfer, error) {
	var txs []*xchain.Transfer

	resp, err := b.getTxnByHash(ctx, txHash)
	if err != nil {
		return txs, err
	}
	fmt.Printf("%v", resp)
	txs, _ = b.parseTxResponse(json.JGet(resp, "tx_response").String())
	err = b.updateTxsStatus(ctx, txs)
	return txs, err
}

func (b *Band) getTxnByHash(ctx context.Context, txHash string) (string, error) {
	return b.nodeClient.Get(ctx, fmt.Sprintf("/cosmos/tx/v1beta1/txs/%s", txHash), nil)
}

func (b *Band) ScanTransfers(ctx context.Context, cursor *xchain.Cursor) ([]*xchain.Transfer, error) {
	latestHeight, err := b.GetLatestHeight(ctx)
	if err != nil {
		return nil, err
	}
	var txs []*xchain.Transfer
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

func (b *Band) scanTxnByBlock(ctx context.Context, height int64) ([]*xchain.Transfer, error) {
	var txs []*xchain.Transfer
	resp, err := b.getTxs(ctx, height, 0, 100)
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

func (b *Band) getTxs(ctx context.Context, height, offset, limit int64) (string, error) {
	params := map[string]string{
		"events":            fmt.Sprintf("tx.height=%d", height),
		"pagination.offset": strconv.FormatInt(offset, 10),
		"pagination.limit":  strconv.FormatInt(limit, 10),
	}

	return b.nodeClient.Get(ctx, "/cosmos/tx/v1beta1/txs", params)
}

func (b *Band) NewAccount(ctx context.Context, label xchain.WalletLabel) (*xchain.Account, error) {
	_ = ctx
	account := &xchain.Account{}
	if label == xchain.WalletLabelDeposit {
		account = &xchain.Account{
			Address: xchain.Address(b.Config.DepositAddress),
			Memo:    xchain.Memo(rand.DigitalMemo()),
		}
	} else if label == xchain.WalletLabelHot {
		account = &xchain.Account{
			Address: xchain.Address(b.Config.HotAddress),
		}
	}
	return account, nil
}

func (b *Band) GetWalletBalance(ctx context.Context, walletLabel xchain.WalletLabel, identity xchain.Identity) (xchain.CoinValue, error) {
	_ = ctx
	_ = walletLabel
	_ = identity
	return xchain.NewCoinValue(string(identity), decimal.Zero()), nil
}

func (b *Band) GetAccountBalance(ctx context.Context, address xchain.Address, identity xchain.Identity) (xchain.CoinValue, error) {
	balance := xchain.CoinValue{Identity: identity, Amount: decimal.Zero()}
	resp, err := b.getAddressBalance(ctx, address)
	if err != nil {
		return balance, err
	}
	for _, balances := range json.JGet(resp, "balances").Array() {
		if json.JGet(balances.String(), "denom").String() == string(identity) {
			var amount decimal.Decimal
			amount, err = decimal.NewFromString(json.JGet(balances.String(), "amount").String())
			if err == nil {
				balance.Amount = balance.Amount.Add(amount)
			}
		}
	}
	return balance, nil
}

func (b *Band) getAddressBalance(ctx context.Context, address xchain.Address) (string, error) {
	url := fmt.Sprintf("/cosmos/bank/v1beta1/balances/%s", address)
	return b.nodeClient.Get(ctx, url, nil)
}

func (b *Band) EstimateFee(ctx context.Context, reqData *xchain.TransferCommand) (*xchain.Fee, error) {
	_ = ctx
	_ = reqData
	return &xchain.Fee{}, nil
}

func (b *Band) Transfer(ctx context.Context, reqData *xchain.TransferCommand) (*xchain.Receipt, error) {
	_ = ctx
	_ = reqData
	return &xchain.Receipt{}, nil
}

func (b *Band) parseTxResponse(txResponse string) ([]*xchain.Transfer, error) {

	var txs []*xchain.Transfer

	status := xchain.TxnPending
	if code := json.JGet(txResponse, "code").Int(); code != 0 {
		status = xchain.TxnFailed
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
			txn := &xchain.Transfer{
				Chain:  b.Code,
				TxHash: txHash,
				VOut:   0,
				Sender: sender,
				Recipient: &xchain.Account{
					Address: xchain.Address(receiver),
					Memo:    xchain.Memo(memo),
				},
				CoinValue: xchain.CoinValue{
					Identity: xchain.Identity(identity),
					Amount:   decimal.NewFromInt(amountInt),
				},
				Status:        status,
				Confirmations: 0,
				Height:        height,
			}
			txs = append(txs, txn)
		}
	}
	return txs, nil
}

func (b *Band) updateTxsStatus(ctx context.Context, txs []*xchain.Transfer) error {
	latestHeight, _ := b.GetLatestHeight(ctx)
	for _, txn := range txs {
		if txn.Status == xchain.TxnFailed {
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
		txn.Confirmations = confirm
		if confirm >= b.Config.IrreversibleBlock {
			txn.Status = xchain.TxnSuccess
		}
	}
	return nil
}
