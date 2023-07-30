package waxp

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/xlalon/golee/internal/xchain"
	"github.com/xlalon/golee/internal/xchain/conf"
	"github.com/xlalon/golee/pkg/json"
	"github.com/xlalon/golee/pkg/math/decimal"
	"github.com/xlalon/golee/pkg/math/rand"
	"github.com/xlalon/golee/pkg/net/http/client"
)

type Waxp struct {
	*xchain.X

	nodeClient   client.Client
	mainContract string
}

func New(conf *conf.ChainConfig) *Waxp {
	return &Waxp{
		X: &xchain.X{
			Code:   "WAX",
			Config: conf,
		},

		nodeClient: client.NewRestClient(conf.NodeUrl),

		mainContract: "eosio.token",
	}
}

func (w *Waxp) GetLatestHeight(ctx context.Context) (int64, error) {
	resp, err := w.getLatestHeight(ctx)
	if err != nil {
		return -1, err
	}
	height := json.JGet(resp, "last_irreversible_block_num")
	return height.Int(), nil
}

func (w *Waxp) getLatestHeight(ctx context.Context) (string, error) {
	return w.nodeClient.Post(ctx, "/v1/chain/get_info", nil)
}

func (w *Waxp) GetTransfersByHash(ctx context.Context, txHash string) ([]*xchain.Transfer, error) {
	var txs []*xchain.Transfer
	resp, err := w.getTransfersByHash(ctx, txHash)
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

func (w *Waxp) getTransfersByHash(ctx context.Context, txHash string) (string, error) {
	params := map[string]string{"id": txHash}
	return w.nodeClient.Get(ctx, "/v2/history/get_transaction", params)
}

func (w *Waxp) ScanTransfers(ctx context.Context, cursor *xchain.Cursor) ([]*xchain.Transfer, error) {
	txs, err := w.scanTxnByAccount(ctx, cursor)
	if err != nil {
		return nil, err
	}
	return txs, nil
}

func (w *Waxp) scanTxnByAccount(ctx context.Context, cursor *xchain.Cursor) ([]*xchain.Transfer, error) {
	var txs []*xchain.Transfer
	resp, err := w.getActions(ctx, cursor.AccountCursor.Address, 1, 100)
	if err != nil {
		return txs, err
	}
	actionsArray := json.JGet(resp, "actions").Array()
	for _, action := range actionsArray {
		if txn, errParsed := w.parseAct(action.String()); txn != nil && errParsed == nil && txn.Recipient.Address == cursor.AccountCursor.Address {
			txs = append(txs, txn)
		}
	}
	cursor.AccountCursor.Index += int64(len(actionsArray))

	return txs, nil
}

func (w *Waxp) getActions(ctx context.Context, address xchain.Address, offset, limit int64) (string, error) {
	params := map[string]string{
		"account": string(address),
		"skip":    strconv.FormatInt(offset, 10),
		"limit":   strconv.FormatInt(limit, 10),
	}
	return w.nodeClient.Get(ctx, "/v2/history/get_actions", params)
}

func (w *Waxp) NewAccount(ctx context.Context, walletLabel xchain.WalletLabel) (*xchain.Account, error) {
	_ = ctx
	account := &xchain.Account{}
	if walletLabel == xchain.WalletLabelDeposit {
		account = &xchain.Account{
			Address: xchain.Address(w.Config.DepositAddress),
			Memo:    xchain.Memo(rand.DigitalMemo()),
		}
	} else if walletLabel == xchain.WalletLabelHot {
		account = &xchain.Account{
			Address: xchain.Address(w.Config.HotAddress),
		}
	}
	return account, nil
}

func (w *Waxp) GetWalletBalance(ctx context.Context, walletLabel xchain.WalletLabel, identity xchain.Identity) (xchain.CoinValue, error) {
	_ = ctx
	_ = walletLabel
	_ = identity
	return xchain.CoinValue{Identity: identity, Amount: decimal.Zero()}, nil

}

func (w *Waxp) GetAccountBalance(ctx context.Context, address xchain.Address, identity xchain.Identity) (xchain.CoinValue, error) {
	balance := xchain.CoinValue{Identity: identity, Amount: decimal.Zero()}
	resp, err := w.getAddressBalance(ctx, address, identity)
	if err != nil {
		return balance, err
	}
	for _, _balance := range json.JParse(resp).Array() {
		if balancesInfo := strings.Split(_balance.String(), string(identity)); len(balancesInfo) > 1 {
			amount, err := decimal.NewFromString(strings.TrimSpace(balancesInfo[0]))
			if err == nil {
				balance.Amount = balance.Amount.Add(amount)
			}
		}
	}
	return balance, nil

}

func (w *Waxp) getAddressBalance(ctx context.Context, address xchain.Address, identity xchain.Identity) (string, error) {
	params := map[string]interface{}{
		"account": address,
		"code":    w.mainContract,
		"symbol":  identity,
	}
	return w.nodeClient.Post(ctx, "/v1/chain/get_currency_balance", params)
}

func (w *Waxp) EstimateFee(ctx context.Context, reqData *xchain.TransferCommand) (*xchain.Fee, error) {
	_ = ctx
	_ = reqData
	return &xchain.Fee{}, nil
}

func (w *Waxp) Transfer(ctx context.Context, reqData *xchain.TransferCommand) (*xchain.Receipt, error) {
	_ = ctx
	_ = reqData
	return &xchain.Receipt{}, nil
}

func (w *Waxp) parseTxResponse(txResponse string) ([]*xchain.Transfer, error) {

	var txs []*xchain.Transfer

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

func (w *Waxp) parseAct(action string) (*xchain.Transfer, error) {

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

	txn := &xchain.Transfer{
		Chain:  w.Code,
		TxHash: txHash,
		VOut:   0,
		Sender: sender,
		Recipient: &xchain.Account{
			Address: xchain.Address(receiver),
			Memo:    xchain.Memo(memo),
		},
		CoinValue: xchain.CoinValue{
			Identity: xchain.Identity(identity),
			Amount:   amount,
		},
		Height:        height,
		Confirmations: 0,
		Status:        xchain.TxnPending,
	}

	return txn, nil
}

func (w *Waxp) updateTxsStatus(ctx context.Context, txs []*xchain.Transfer) error {
	latestHeight, _ := w.GetLatestHeight(ctx)
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
		if confirm >= w.Config.IrreversibleBlock {
			txn.Status = xchain.TxnSuccess
		}
	}
	return nil
}
