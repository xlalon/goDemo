package model

import (
	"gorm.io/gorm"

	"github.com/xlalon/golee/pkg/math/decimal"
)

type Deposit struct {
	gorm.Model

	Chain     string
	Asset     string
	TxHash    string
	Sender    string
	Receiver  string
	Memo      string
	Identity  string
	Amount    decimal.Decimal
	AmountRaw decimal.Decimal
	VOut      int64
	Status    string
	Comment   string

	Version int64
}
