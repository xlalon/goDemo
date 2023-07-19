package xchain

import (
	"fmt"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type Identity string

type Coin struct {
	Identity Identity        `json:"identity"`
	Amount   decimal.Decimal `json:"amount"`
}

func (c Coin) String() string {
	return fmt.Sprintf("%s %v", c.Identity, c.Amount)
}
