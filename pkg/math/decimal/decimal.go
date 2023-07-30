package decimal

import (
	xdecimal "github.com/shopspring/decimal"
	"math/big"
)

type Decimal struct {
	xdecimal.Decimal
}

func NewFromString(value string) (Decimal, error) {
	result, err := xdecimal.NewFromString(value)
	return Decimal{result}, err
}

func Zero() Decimal {
	return Decimal{xdecimal.NewFromInt(0)}
}

func NewFromInt(value int64) Decimal {
	return Decimal{xdecimal.NewFromInt(value)}
}

func NewFromBigInt(value *big.Int, exp int32) Decimal {
	return Decimal{xdecimal.NewFromBigInt(value, exp)}
}

func NewFromFloat(value float64) Decimal {
	return Decimal{xdecimal.NewFromFloat(value)}
}

func (d Decimal) Truncate(precision int32) Decimal {
	return Decimal{d.Decimal.Truncate(precision)}
}

func (d Decimal) String() string {
	return d.Decimal.String()
}

func (d Decimal) Add(d2 Decimal) Decimal {
	return Decimal{d.Decimal.Add(d2.Decimal)}
}

func (d Decimal) Sub(d2 Decimal) Decimal {
	return Decimal{d.Decimal.Sub(d2.Decimal)}
}

func (d Decimal) Mul(d2 Decimal) Decimal {
	return Decimal{d.Decimal.Mul(d2.Decimal)}
}

func (d Decimal) Div(d2 Decimal) Decimal {
	return Decimal{d.Decimal.Div(d2.Decimal)}
}

func (d Decimal) Pow(d2 Decimal) Decimal {
	return Decimal{d.Decimal.Pow(d2.Decimal)}
}

func (d Decimal) IsZero() bool {
	return d.Decimal.IsZero()
}

func (d Decimal) Equal(d2 Decimal) bool {
	return d.Decimal.Equal(d2.Decimal)
}

func (d Decimal) GreaterThan(d2 Decimal) bool {
	return d.Decimal.GreaterThan(d2.Decimal)
}

func (d Decimal) GreaterThanZero() bool {
	return d.Decimal.GreaterThan(xdecimal.NewFromInt(0))
}

func (d Decimal) GreaterThanOrEqual(d2 Decimal) bool {
	return d.Decimal.GreaterThanOrEqual(d2.Decimal)
}

func (d Decimal) GreaterThanOrEqualZero() bool {
	return d.Decimal.GreaterThanOrEqual(xdecimal.NewFromInt(0))
}

func (d Decimal) LessThan(d2 Decimal) bool {
	return d.Decimal.LessThan(d2.Decimal)
}

func (d Decimal) LessThanOrEqual(d2 Decimal) bool {
	return d.Decimal.LessThanOrEqual(d2.Decimal)
}
