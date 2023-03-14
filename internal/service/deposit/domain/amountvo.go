package domain

import (
	"github.com/xlalon/golee/pkg/math/decimal"
)

type AmountVO struct {
	identity   string
	amountRaw  decimal.Decimal
	precession int64
	amount     decimal.Decimal
}

func NewAmountVO(identity string, amountRaw decimal.Decimal, precession int64, amount decimal.Decimal) AmountVO {
	return AmountVO{
		identity:   identity,
		amountRaw:  amountRaw,
		precession: precession,
		amount:     amount,
	}
}

func (a AmountVO) GetIdentity() string {
	return a.identity
}

func (a AmountVO) GetAmountRaw() decimal.Decimal {
	return a.amountRaw
}

func (a AmountVO) GetPrecession() int64 {
	return a.precession
}

func (a AmountVO) GetAmount() decimal.Decimal {
	return a.amount
}

func (a AmountVO) ToAmountRaw() decimal.Decimal {
	return a.amount.Mul(decimal.NewFromInt(10).Pow(decimal.NewFromInt(a.precession)))
}

func (a AmountVO) ToAmount() decimal.Decimal {
	return a.amountRaw.Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(a.precession)))
}
