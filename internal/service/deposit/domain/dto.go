package domain

import (
	"github.com/xlalon/golee/pkg/math/decimal"
)

type DepositDTO struct {
	Id         int64           `json:"id"`
	Chain      string          `json:"chain"`
	Asset      string          `json:"asset"`
	TxHash     string          `json:"tx_hash"`
	Sender     string          `json:"sender"`
	Receiver   string          `json:"receiver"`
	Memo       string          `json:"memo"`
	Identity   string          `json:"identity"`
	AmountRaw  decimal.Decimal `json:"amount_raw"`
	Precession int64           `json:"-"`
	Amount     decimal.Decimal `json:"amount"`
	VOut       int64           `json:"v_out"`
	Status     Status          `json:"status"`
}
