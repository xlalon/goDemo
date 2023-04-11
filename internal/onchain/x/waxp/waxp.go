package waxp

import (
	"context"
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

	nodeClient   client.Client
	mainContract string
}

func New(conf *conf.ChainConfig) *Waxp {
	return &Waxp{
		Chain: &onchain.Chain{
			Code:   "WAX",
			Config: conf,
		},

		nodeClient: client.NewRestClient(conf.NodeUrl),

		mainContract: "eosio.token",
	}
}

func (w *Waxp) GetLatestHeight(ctx context.Context) (int64, error) {
	resp, err := w.nodeClient.Post(ctx, "/v1/chain/get_info", nil)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	height := json.JGet(resp, "last_irreversible_block_num")
	return height.Int(), nil
}

func (w *Waxp) GetTxnByHash(ctx context.Context, txHash string) ([]*onchain.Transaction, error) {
	var txs []*onchain.Transaction
	json_ := map[string]interface{}{"id": txHash}
	resp, err := w.nodeClient.Post(ctx, "/v1/history/get_transaction", json_)
	if err != nil {
		fmt.Println(err)
		return txs, err
	}

	txs, err = w.parseTxResponse(resp)
	if err != nil {
		return txs, err
	}

	err = w.updateTxsStatus(ctx, txs)

	return txs, err
}

func (w *Waxp) ScanTxn(ctx context.Context, cursor *onchain.Cursor) ([]*onchain.Transaction, error) {
	txs, err := w.scanTxnByAccount(ctx, cursor)
	if err != nil {
		return nil, err
	}

	cursor.Index += 10

	return txs, nil
}

func (w *Waxp) scanTxnByAccount(ctx context.Context, cursor *onchain.Cursor) ([]*onchain.Transaction, error) {

	var txs []*onchain.Transaction

	data := map[string]interface{}{
		"account_name": cursor.Account.Address,
		"pos":          cursor.Index,
		"offset":       -10,
	}
	resp, err := w.nodeClient.Post(ctx, "/v1/history/get_actions", data)
	if err != nil {
		fmt.Println(err)
		return txs, err
	}

	actionsArray := json.JGet(resp, "actions").Array()
	for _, actionRaw := range actionsArray {
		action := actionRaw.String()
		if json.JGet(action, "action_trace.action_ordinal").Int() < 3 {
			continue
		}
		height := json.JGet(action, "block_num").Int()
		txHash := json.JGet(action, "action_trace.trx_id").String()
		if txn, errParsed := w.parseAct(json.JGet(action, "action_trace.act").String()); txn != nil && errParsed == nil && txn.CoinValue.Amount.GreaterThan(decimal.NewFromInt(0)) && txn.Receiver.Address == cursor.Account.Address {
			txn.TxnId.TxHash = txHash
			txn.Height = height
			txs = append(txs, txn)
		}
	}

	cursor.Index += int64(len(actionsArray))

	return txs, nil
}

func (w *Waxp) NewAccount(ctx context.Context, label onchain.Label) (*onchain.Account, error) {
	_ = ctx
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

func (w *Waxp) GetAccount(ctx context.Context, address string) (*onchain.Account, error) {
	_ = ctx
	_ = address
	return &onchain.Account{}, nil
}

func (w *Waxp) EstimateFee(ctx context.Context, reqData *onchain.TransferCommand) (*onchain.Fee, error) {
	_ = ctx
	_ = reqData
	return &onchain.Fee{}, nil
}

func (w *Waxp) Transfer(ctx context.Context, reqData *onchain.TransferCommand) (*onchain.Receipt, error) {
	_ = ctx
	_ = reqData
	return &onchain.Receipt{}, nil
}

func (w *Waxp) parseTxResponse(txResponse string) ([]*onchain.Transaction, error) {

	var txs []*onchain.Transaction

	receipt := json.JGet(txResponse, "trx.receipt").String()
	if status := json.JGet(receipt, "status"); status.String() != "executed" {
		return txs, fmt.Errorf("tx status not right, expect %s, got %s", "executed", status.String())
	}
	height := json.JGet(txResponse, "block_num").Int()
	txHash := json.JGet(txResponse, "id").String()
	actions := json.JGet(txResponse, "trx.trx.actions")
	for _, actionRaw := range actions.Array() {
		action := actionRaw.String()
		if txn, _ := w.parseAct(action); txn != nil && txn.CoinValue.Amount.GreaterThan(decimal.NewFromInt(0)) {
			txn.TxnId.TxHash = txHash
			txn.Height = height
			txs = append(txs, txn)
		}
	}
	return txs, nil
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

	txn := &onchain.Transaction{
		TxnId: onchain.TxnId{
			Chain:  w.Code,
			TxHash: "",
			VOut:   0,
		},
		Receiver: onchain.Receiver{
			Address: receiver,
			Memo:    memo,
		},
		CoinValue: onchain.CoinValue{
			Identity: amountInfo[1],
			Amount:   amount,
		},
		Sender: sender,
		Status: onchain.TxnStatus{Result: onchain.TxnPending},
	}

	return txn, nil
}

func (w *Waxp) updateTxsStatus(ctx context.Context, txs []*onchain.Transaction) error {
	latestHeight, _ := w.GetLatestHeight(ctx)
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
		if confirm >= w.Config.IrreversibleBlock {
			txn.Status.Result = onchain.TxnSuccess
		}
	}
	return nil
}
