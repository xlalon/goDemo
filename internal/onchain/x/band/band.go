package band

import (
	"errors"
	"fmt"
	"net/url"
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

	nodeClient *client.RestfulClient

	maxScanHeightRange int64
}

func New(conf *conf.ChainConfig) *Band {
	return &Band{
		Chain: &onchain.Chain{
			Code:   "BAND",
			Config: conf,
		},

		nodeClient: client.NewRestfulClient(
			&client.Config{
				BaseUrl: conf.NodeUrl,
				Headers: url.Values{"Content-Type": []string{"application/json"}},
				Timeout: 30,
			}),
		maxScanHeightRange: 2,
	}
}

func (b *Band) GetLatestHeight() (int64, error) {
	resp, err := b.nodeClient.Get("/blocks/latest", nil)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	height := json.JGet(resp, "block.header.height")
	return height.Int(), nil
}

func (b *Band) GetTxnByHash(txHash string) ([]*onchain.Transaction, error) {
	var txs []*onchain.Transaction

	resp, err := b.nodeClient.Get(fmt.Sprintf("/cosmos/tx/v1beta1/txs/%s", txHash), nil)
	if err != nil {
		fmt.Println(err)
		return txs, err
	}

	txs, _ = b.parseTxResponse(json.JGet(resp, "tx_response").String())

	err = b.updateTxsStatus(txs)

	return txs, err
}

func (b *Band) ScanTxn(cursor *onchain.Cursor) ([]*onchain.Transaction, error) {
	latestHeight, err := b.GetLatestHeight()
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
		_txs, errScan := b.scanTxnByBlock(height)
		if errScan != nil {
			return nil, errScan
		}
		txs = append(txs, _txs...)
	}

	cursor.Height = maxHeight

	return txs, nil
}

func (b *Band) scanTxnByBlock(height int64) ([]*onchain.Transaction, error) {

	var txs []*onchain.Transaction

	params := make(url.Values)
	params.Set("events", fmt.Sprintf("tx.height=%d", height))
	params.Set("pagination.offset", strconv.Itoa(1))
	params.Set("pagination.limit", strconv.Itoa(100))

	resp, err := b.nodeClient.Get("/cosmos/tx/v1beta1/txs", params)
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

func (b *Band) NewAccount(label onchain.Label) (*onchain.Account, error) {
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

func (b *Band) GetAccount(address string) (*onchain.Account, error) {

	// balance info
	reqBalanceUrl := fmt.Sprintf("/cosmos/bank/v1beta1/balances/%s", address)
	respBalance, err := b.nodeClient.Get(reqBalanceUrl, nil)
	if err != nil {
		return nil, err
	}
	var balances []*onchain.Balance
	for _, balanceInfo := range json.JGet(respBalance, "balances").Array() {
		balanceInfoStr := balanceInfo.String()
		amount := decimal.Decimal{}
		amount, err = decimal.NewFromString(json.JGet(balanceInfoStr, "amount").String())
		if err != nil {
			return nil, err
		}
		balances = append(balances, &onchain.Balance{
			Identity: json.JGet(balanceInfoStr, "denom").String(),
			Amount:   amount,
		})
	}
	// sequence info
	reqAccountUrl := fmt.Sprintf("/cosmos/auth/v1beta1/accounts/%s", address)
	respAccount, err := b.nodeClient.Get(reqAccountUrl, nil)
	if err != nil {
		return nil, err
	}
	sequence := json.JGet(respAccount, "account.sequence").Int()
	if err != nil {
		return nil, err
	}

	account := &onchain.Account{
		Chain:    b.Code,
		Address:  address,
		Label:    onchain.AccountUnknown,
		Sequence: sequence,
		Balance:  balances,
	}

	return account, nil
}

func (b *Band) EstimateFee(reqData *onchain.TransferDTO) (*onchain.Fee, error) {
	_ = reqData
	return &onchain.Fee{}, nil
}

func (b *Band) Transfer(reqData *onchain.TransferDTO) (*onchain.Receipt, error) {
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

func (b *Band) updateTxsStatus(txs []*onchain.Transaction) error {
	latestHeight, _ := b.GetLatestHeight()
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
