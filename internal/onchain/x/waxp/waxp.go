package waxp

import (
	"context"
	"errors"
	"fmt"
	"strconv"
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
	params := map[string]string{"id": txHash}
	resp, err := w.nodeClient.Get(ctx, "/v2/history/get_transaction", params)
	if err != nil {
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

	params := map[string]string{
		"account": cursor.Account.Address,
		"skip":    strconv.Itoa(0),
		"limit":   strconv.Itoa(10),
	}
	resp, err := w.nodeClient.Get(ctx, "/v2/history/get_actions", params)
	if err != nil {
		return txs, err
	}

	actionsArray := json.JGet(resp, "actions").Array()
	for _, action := range actionsArray {
		if txn, errParsed := w.parseAct(action.String()); txn != nil && errParsed == nil && txn.Receiver.Address == cursor.Account.Address {
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
	account := &onchain.Account{
		Chain:   w.Code,
		Address: address,
		Label:   onchain.AccountUnknown,
	}
	return account, nil
}

func (w *Waxp) GetBalance(ctx context.Context, account *onchain.Account, identity string) (decimal.Decimal, error) {
	zero := decimal.Zero()
	params := map[string]interface{}{
		"account": account.Address,
		"code":    w.mainContract,
		"symbol":  identity,
	}
	resp, err := w.nodeClient.Post(ctx, "/v1/chain/get_currency_balance", params)
	if err != nil {
		return zero, err
	}
	for _, balance := range json.JParse(resp).Array() {
		if balancesInfo := strings.Split(balance.String(), identity); len(balancesInfo) > 1 {
			return decimal.NewFromString(strings.TrimSpace(balancesInfo[0]))
		}
	}
	return zero, errors.New("identity not found")

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

	executed := json.JGet(txResponse, "executed").Bool()
	if !executed {
		return txs, fmt.Errorf("tx not executed")
	}
	actions := json.JGet(txResponse, "actions")
	for _, actionRaw := range actions.Array() {
		if txn, _ := w.parseAct(actionRaw.String()); txn != nil && txn.CoinValue.Amount.GreaterThan(decimal.NewFromInt(0)) {
			txs = append(txs, txn)
		}
	}
	return txs, nil
}

func (w *Waxp) parseAct(action string) (*onchain.Transaction, error) {

	if json.JGet(action, "act.account").String() != w.mainContract || json.JGet(action, "act.name").String() != "transfer" {
		return nil, errors.New("parse failed")
	}

	txHash := json.JGet(action, "trx_id").String()
	height := json.JGet(action, "block_num").Int()

	data := json.JGet(action, "act.data").String()

	identity := json.JGet(data, "symbol").String()
	amount := decimal.NewFromFloat(json.JGet(data, "amount").Float())
	if !amount.GreaterThanZero() {
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
			TxHash: txHash,
			VOut:   0,
		},
		Receiver: onchain.Receiver{
			Address: receiver,
			Memo:    memo,
		},
		CoinValue: onchain.CoinValue{
			Identity: identity,
			Amount:   amount,
		},
		Sender: sender,
		Height: height,
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
