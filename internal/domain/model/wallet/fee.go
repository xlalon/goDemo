package wallet

import "github.com/xlalon/golee/pkg/math/decimal"

type Fee struct {
	Identity string          `json:"identity"`
	Amount   decimal.Decimal `json:"amount"`
	Gas      int64           `json:"gas"`
	GasPrice decimal.Decimal `json:"gas_price"`
}
