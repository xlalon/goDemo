package band

import (
	"github.com/xlalon/golee/common/math/decimal"
	"github.com/xlalon/golee/x/chain"
)

type Wallet struct {
	chain    string
	label    string
	accounts []*Account
	balance  []*chain.Crypto
}

func NewWallet(label string) *Wallet {
	return &Wallet{
		chain:   "BAND",
		label:   label,
		balance: nil,
	}
}

func (w *Wallet) Chain() string {
	return w.chain
}

func (w *Wallet) Label() string {
	return w.label
}

func (w *Wallet) Accounts() []*Account {
	return nil
}

func (w *Wallet) Balance(crypto *chain.Crypto) decimal.Decimal {
	return decimal.Zero()
}

func (w *Wallet) Balances() []*chain.Balance {
	return nil
}

type Account struct {
}
