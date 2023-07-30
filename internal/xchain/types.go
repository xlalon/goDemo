package xchain

import (
	"fmt"
	"strings"

	"github.com/xlalon/golee/pkg/math/decimal"
)

type Chain string

func (c Chain) Normalize() Chain {
	return Chain(strings.ToUpper(string(c)))
}

type Identity string

type CoinValue struct {
	Identity Identity        `json:"identity"`
	Amount   decimal.Decimal `json:"amount"`
}

func NewCoinValue(identity string, amount decimal.Decimal) CoinValue {
	return CoinValue{Identity: Identity(identity), Amount: amount}
}

func (cv CoinValue) String() string {
	return fmt.Sprintf("%v %s", cv.Amount, cv.Identity)
}

type WalletLabel string

const (
	WalletLabelDeposit WalletLabel = "DEPOSIT"
	WalletLabelHot     WalletLabel = "HOT"
)

type Wallet struct {
	Chain    Chain       `json:"chain"`
	Label    WalletLabel `json:"label"`
	Accounts []*Account  `json:"accounts"`
}

type Address string
type Memo string

type Account struct {
	Address    Address `json:"address"`
	Memo       Memo    `json:"memo"`
	privateKey []byte
	publicKey  []byte
	index      int64
}

type TxnStatus string

const (
	TxnFailed  TxnStatus = "FAILED"
	TxnPending TxnStatus = "PENDING"
	TxnSuccess TxnStatus = "SUCCESS"
)

type Transfer struct {
	Chain         Chain     `json:"chain"`
	TxHash        string    `json:"tx_hash"`
	VOut          int64     `json:"v_out"`
	Sender        string    `json:"sender"`
	Recipient     *Account  `json:"recipient"`
	CoinValue     CoinValue `json:"coin"`
	Status        TxnStatus `json:"status"`
	Confirmations int64     `json:"confirmations"`
	Height        int64     `json:"height"`
	Timestamp     int64     `json:"timestamp"`
	Comment       string    `json:"comment"`
}

type Fee struct {
	CoinValue CoinValue `json:"coin"`
	GasPrice  int64     `json:"gas_price"`
	GasLimit  int64     `json:"gas_limit"`
	GasUsed   int64     `json:"gas_used"`
}

func NewFee(identity string, amount decimal.Decimal, gasPrice, gasLimit, gasUsed int64) *Fee {
	return &Fee{
		CoinValue: CoinValue{
			Identity: Identity(identity),
			Amount:   amount,
		},
		GasPrice: gasPrice,
		GasLimit: gasLimit,
		GasUsed:  gasUsed,
	}
}

type Receipt struct {
	TxHash string    `json:"tx_hash"`
	Fee    *Fee      `json:"fee"`
	Status TxnStatus `json:"status"`
	ErrLog string    `json:"err_log"`
}

type AccountCursor struct {
	Address Address `json:"address"`
	TxHash  string  `json:"tx_hash"`
	Index   int64   `json:"index"`
}

type WalletCursor struct {
	WalletLabel WalletLabel `json:"wallet_label"`
	TxHash      string      `json:"tx_hash"`
	Index       int64       `json:"index"`
}

type Cursor struct {
	// scan by height
	Height int64 `json:"height"`
	// scan by account
	AccountCursor *AccountCursor `json:"account_cursor"`
	// scan by wallet
	WalletCursor *WalletCursor `json:"wallet_cursor"`
}

type TransferCommand struct {
	Sender    *Wallet                `json:"sender"`
	Recipient *Account               `json:"receiver"`
	CoinValue CoinValue              `json:"coin_value"`
	Fee       *Fee                   `json:"fee"`
	Extra     map[string]interface{} `json:"extra"`
}

func NewTransferCommand(chain, walletLabel, sender, recipient, memo, identity string, amount decimal.Decimal, fee *Fee, extra map[string]interface{}) *TransferCommand {

	return &TransferCommand{
		Sender: &Wallet{
			Chain: Chain(chain),
			Label: WalletLabel(walletLabel),
			Accounts: []*Account{
				{Address: Address(sender)},
			},
		},
		Recipient: &Account{
			Address: Address(recipient),
			Memo:    Memo(memo),
		},
		CoinValue: CoinValue{
			Identity: Identity(identity),
			Amount:   amount,
		},
		Fee:   fee,
		Extra: extra,
	}
}
