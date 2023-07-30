package wallet

import "github.com/xlalon/golee/pkg/math/decimal"

type AccountDTO struct {
	Id       int64         `json:"id"`
	Chain    string        `json:"chain"`
	Address  string        `json:"address"`
	Memo     string        `json:"memo"`
	Status   string        `json:"status"`
	Balances []*BalanceDTO `json:"balances"`
}

type BalanceDTO struct {
	Amount decimal.Decimal `json:"amount"`
	Asset  string          `json:"asset"`
}
