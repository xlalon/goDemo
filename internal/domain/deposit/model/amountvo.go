package model

import (
	"github.com/xlalon/golee/pkg/ecode"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type AmountVO struct {
	identity   string
	amountRaw  decimal.Decimal
	precession int64
	amount     decimal.Decimal
}

func NewAmountVO(identity string, amountRaw decimal.Decimal, precession int64, amount decimal.Decimal) *AmountVO {
	amountVO := &AmountVO{}
	if err := amountVO.setIdentity(identity); err != nil {
		return nil
	}
	if err := amountVO.setAmountRaw(amountRaw); err != nil {
		return nil
	}
	if err := amountVO.setPrecession(precession); err != nil {
		return nil
	}
	if err := amountVO.setAmount(amount); err != nil {
		return nil
	}
	return amountVO
}

func (a *AmountVO) Identity() string {
	return a.identity
}

func (a *AmountVO) setIdentity(identity string) error {
	if a.identity != "" {
		return ecode.ParameterChangeError
	}
	if identity == "" {
		return ecode.ParameterInvalidError
	}
	a.identity = identity
	return nil
}

func (a *AmountVO) AmountRaw() decimal.Decimal {
	return a.amountRaw
}

func (a *AmountVO) setAmountRaw(amountRaw decimal.Decimal) error {
	if a.amountRaw.GreaterThanZero() {
		return ecode.ParameterChangeError
	}
	if !amountRaw.GreaterThanOrEqualZero() {
		return ecode.ParameterInvalidError
	}
	a.amountRaw = amountRaw
	return nil
}

func (a *AmountVO) Precession() int64 {
	return a.precession
}

func (a *AmountVO) setPrecession(precession int64) error {
	if a.precession > 0 {
		return ecode.ParameterChangeError
	}
	if precession < 0 {
		return ecode.ParameterInvalidError
	}
	a.precession = precession
	return nil
}

func (a *AmountVO) Amount() decimal.Decimal {
	return a.amount
}

func (a *AmountVO) setAmount(amount decimal.Decimal) error {
	if a.amount.GreaterThanZero() {
		return ecode.ParameterChangeError
	}
	if !amount.GreaterThanOrEqualZero() {
		return ecode.ParameterInvalidError
	}
	a.amount = amount
	return nil
}

func (a *AmountVO) ToAmountRaw() decimal.Decimal {
	return a.amount.Mul(decimal.NewFromInt(10).Pow(decimal.NewFromInt(a.precession)))
}

func (a *AmountVO) ToAmount() decimal.Decimal {
	if a.precession == 0 {
		return a.amountRaw
	}
	return a.amountRaw.Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(a.precession)))
}
