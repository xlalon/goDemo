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
	"github.com/xlalon/golee/pkg/net/http/client"
)

type Band struct {
	*onchain.Chain

	nodeClient *client.RestfulClient
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
	var transfers []*onchain.Transaction

	resp, err := b.nodeClient.Get(fmt.Sprintf("/cosmos/tx/v1beta1/txs/%s", txHash), nil)
	if err != nil {
		fmt.Println(err)
		return transfers, err
	}

	transfers, _ = b.parseTx(json.JGet(resp, "tx_response").String())

	err = b.updateTransfersStatus(transfers)

	return transfers, err
}

func (b *Band) ScanTxn(xxx interface{}) ([]*onchain.Transaction, error) {
	return b.ScanTxnByBlock(xxx)
}

func (b *Band) ScanTxnByBlock(heightOrHash interface{}) ([]*onchain.Transaction, error) {

	var transfers []*onchain.Transaction
	height := heightOrHash.(int64)

	params := make(url.Values)
	params.Set("events", fmt.Sprintf("tx.height=%d", height))
	params.Set("pagination.offset", strconv.Itoa(1))
	params.Set("pagination.limit", strconv.Itoa(100))

	resp, err := b.nodeClient.Get("/cosmos/tx/v1beta1/txs", params)
	if err != nil {
		return transfers, err
	}
	txsResp := json.JGet(resp, "tx_responses")

	for _, txResp := range txsResp.Array() {
		parsedTx, _ := b.parseTx(txResp.String())
		transfers = append(transfers, parsedTx...)
	}
	return transfers, nil
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

func (b *Band) parseTx(txResponse string) ([]*onchain.Transaction, error) {

	var transfers []*onchain.Transaction

	status := onchain.TransactionStatus{Result: onchain.TransactionPending}
	if code := json.JGet(txResponse, "code").Int(); code != 0 {
		status.Result = onchain.TransactionFailed
	}

	tx := json.JGet(txResponse, "tx").String()
	if txType := json.JGet(tx, "@type").String(); txType != "/cosmos.tx.v1beta1.Tx" {
		return transfers, errors.New("tx type not /cosmos.tx.v1beta1.Tx")
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
			transfer := &onchain.Transaction{
				Chain:    b.Code,
				Identity: identity,
				TxHash:   txHash,
				Height:   height,
				Sender:   sender,
				Receiver: receiver,
				Amount:   decimal.NewFromInt(amountInt),
				VOut:     0,
				Memo:     memo,
				Status:   status,
			}
			transfers = append(transfers, transfer)
		}
	}
	return transfers, nil
}

func (b *Band) updateTransfersStatus(transfers []*onchain.Transaction) error {
	latestHeight, _ := b.GetLatestHeight()
	for _, transfer := range transfers {
		if transfer.Status.Result == onchain.TransactionFailed {
			continue
		}
		if transfer.Height == 0 {
			// update transfer height first
			continue
		}
		confirm := latestHeight - transfer.Height
		if confirm <= 0 {
			continue
		}
		transfer.Status.Confirmations = confirm
		if confirm >= b.Config.IrreversibleBlock {
			transfer.Status.Result = onchain.TransactionSuccess
		}
	}
	return nil
}
