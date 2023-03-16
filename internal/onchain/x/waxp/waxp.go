package waxp

import (
	"errors"
	"fmt"
	"strings"

	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/internal/onchain/conf"
	"github.com/xlalon/golee/pkg/json"
	"github.com/xlalon/golee/pkg/math/decimal"
	"github.com/xlalon/golee/pkg/math/rand"
	"github.com/xlalon/golee/pkg/net/http/client"
)

type Waxp struct {
	*onchain.Chain

	nodeClient   *client.RestfulClient
	mainContract string
}

func New(conf *conf.ChainConfig) *Waxp {
	return &Waxp{
		Chain: &onchain.Chain{
			Code:   "WAX",
			Config: conf,
		},

		nodeClient: client.NewRestfulClient(&client.Config{
			BaseUrl: conf.NodeUrl,
		}),

		mainContract: "eosio.token",
	}
}

func (w *Waxp) GetLatestHeight() (int64, error) {
	resp, err := w.nodeClient.Post("/v1/chain/get_info", nil)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	height := json.JGet(resp, "last_irreversible_block_num")
	return height.Int(), nil
}

func (w *Waxp) GetTxnByHash(txHash string) ([]*onchain.Transaction, error) {
	var transfers []*onchain.Transaction
	json_ := map[string]interface{}{"id": txHash}
	resp, err := w.nodeClient.Post("/v1/history/get_transaction", json_)
	if err != nil {
		fmt.Println(err)
		return transfers, err
	}

	transfers, err = w.parseTx(resp)
	if err != nil {
		return transfers, err
	}

	err = w.updateTransfersStatus(transfers)

	return transfers, err
}

func (w *Waxp) ScanTxn(xxx interface{}) ([]*onchain.Transaction, error) {
	return w.ScanTxnByAccount(xxx.(*onchain.Account))
}

func (w *Waxp) ScanTxnByAccount(account *onchain.Account) ([]*onchain.Transaction, error) {

	var transfers []*onchain.Transaction

	data := map[string]interface{}{
		"account_name": account.Address,
		"pos":          account.Sequence,
		"offset":       -10,
	}
	resp, err := w.nodeClient.Post("/v1/history/get_actions", data)
	if err != nil {
		fmt.Println(err)
		return transfers, err
	}

	for _, actionRaw := range json.JGet(resp, "actions").Array() {
		action := actionRaw.String()
		if json.JGet(action, "action_trace.action_ordinal").Int() < 3 {
			continue
		}
		height := json.JGet(action, "block_num").Int()
		txHash := json.JGet(action, "action_trace.trx_id").String()
		if transfer, err := w.parseAct(json.JGet(action, "action_trace.act").String()); transfer != nil && err == nil && transfer.Amount.GreaterThan(decimal.NewFromInt(0)) &&
			transfer.Receiver == account.Address {

			transfer.TxHash = txHash
			transfer.Height = height
			transfers = append(transfers, transfer)
		}
	}
	return transfers, nil
}

func (w *Waxp) NewAccount(label onchain.Label) (*onchain.Account, error) {
	account := &onchain.Account{}
	if label == onchain.AccountDeposit {
		account = &onchain.Account{
			Chain:   w.Code,
			Address: w.Config.DepositAddress,
			Label:   onchain.AccountDeposit,
			Memo:    rand.DigitalMemo(),
		}
	} else if label == onchain.AccountHot {
		account = &onchain.Account{
			Chain:   w.Code,
			Address: w.Config.HotAddress,
			Label:   onchain.AccountHot,
		}
	}
	return account, nil
}

func (w *Waxp) GetAccount(address string) (*onchain.Account, error) {
	_ = address
	return &onchain.Account{}, nil
}

func (w *Waxp) EstimateFee(reqData *onchain.TransferDTO) (*onchain.Fee, error) {
	_ = reqData
	return &onchain.Fee{}, nil
}

func (w *Waxp) Transfer(reqData *onchain.TransferDTO) (*onchain.Receipt, error) {
	_ = reqData
	return &onchain.Receipt{}, nil
}

func (w *Waxp) parseTx(txResponse string) ([]*onchain.Transaction, error) {

	var transfers []*onchain.Transaction

	receipt := json.JGet(txResponse, "trx.receipt").String()
	if status := json.JGet(receipt, "status"); status.String() != "executed" {
		return transfers, fmt.Errorf("tx status not right, expect %s, got %s", "executed", status.String())
	}
	height := json.JGet(txResponse, "block_num").Int()
	txHash := json.JGet(txResponse, "id").String()
	actions := json.JGet(txResponse, "trx.trx.actions")
	for _, actionRaw := range actions.Array() {
		action := actionRaw.String()
		if transfer, _ := w.parseAct(action); transfer != nil && transfer.Amount.GreaterThan(decimal.NewFromInt(0)) {
			transfer.TxHash = txHash
			transfer.Height = height
			transfers = append(transfers, transfer)
		}
	}
	return transfers, nil
}

func (w *Waxp) parseAct(act string) (*onchain.Transaction, error) {

	if json.JGet(act, "account").String() != w.mainContract || json.JGet(act, "name").String() != "transfer" {
		return nil, errors.New("parse failed")
	}

	data := json.JGet(act, "data").String()
	amountInfo := strings.Split(json.JGet(data, "quantity").String(), " ")
	if len(amountInfo) < 2 {
		return nil, errors.New("amount info parse failed")
	}
	amount, _ := decimal.NewFromString(amountInfo[0])
	if amount.LessThanOrEqual(decimal.NewFromInt(0)) {
		return nil, errors.New("amount less than or equal 0")
	}
	receiver := json.JGet(data, "to").String()
	if receiver == "" {
		return nil, errors.New("receiver empty")
	}

	sender := json.JGet(data, "from").String()
	memo := json.JGet(data, "memo").String()

	transfer := &onchain.Transaction{
		Chain:    w.Code,
		Identity: amountInfo[1],
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
		VOut:     0,
		Memo:     memo,
		Status:   onchain.TransactionStatus{Result: onchain.TransactionPending},
	}

	return transfer, nil
}

func (w *Waxp) updateTransfersStatus(transfers []*onchain.Transaction) error {
	latestHeight, _ := w.GetLatestHeight()
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
		if confirm >= w.Config.IrreversibleBlock {
			transfer.Status.Result = onchain.TransactionSuccess
		}
	}
	return nil
}
