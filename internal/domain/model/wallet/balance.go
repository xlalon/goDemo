package wallet

import "github.com/xlalon/golee/pkg/math/decimal"

type Balance struct {
	Asset      string
	Identity   string
	Precession int64
	Amount     decimal.Decimal
	AmountRaw  decimal.Decimal
}

func (b *Balance) ToBalanceDTO() *BalanceDTO {
	return &BalanceDTO{
		Asset:      b.Asset,
		Identity:   b.Identity,
		Precession: b.Precession,
		Amount:     b.Amount,
		AmountRaw:  b.AmountRaw,
	}
}
