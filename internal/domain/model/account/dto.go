package account

import "github.com/xlalon/golee/pkg/math/decimal"

type AccountDTO struct {
	Id       int64         `json:"id"`
	Chain    string        `json:"chain"`
	Address  string        `json:"address"`
	Label    string        `json:"label"`
	Memo     string        `json:"memo"`
	Status   string        `json:"status"`
	Sequence int64         `json:"sequence"`
	Balances []*BalanceDTO `json:"balances"`
}

type BalanceDTO struct {
	Asset      string          `json:"asset"`
	Identity   string          `json:"identity"`
	Precession int64           `json:"precession"`
	Amount     decimal.Decimal `json:"amount"`
	AmountRaw  decimal.Decimal `json:"amount_raw"`
}
