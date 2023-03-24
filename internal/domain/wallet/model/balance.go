package model

import "github.com/xlalon/golee/pkg/math/decimal"

type Balance struct {
	Identity string          `json:"identity"`
	Amount   decimal.Decimal `json:"amount"`
}
