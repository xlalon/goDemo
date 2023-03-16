package onchain

import "github.com/xlalon/golee/pkg/math/decimal"

type Label string

const (
	AccountDeposit  Label = "DEPOSIT"
	AccountHot            = "HOT"
	AccountExternal       = "EXTERNAL"
	AccountUnknown        = "UNKNOWN"
)

type Account struct {
	Chain    Code   `json:"chain"`
	Address  string `json:"address"`
	Label    Label  `json:"label"`
	Memo     string `json:"memo"`
	Sequence int64  `json:"sequence"`

	Balance []*Balance `json:"balance"`
	// uncommon information
	Extra interface{}
}

type Balance struct {
	Identity string          `json:"identity"`
	Amount   decimal.Decimal `json:"amount"`
}
