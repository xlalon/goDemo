package deposit

import (
	"github.com/xlalon/golee/pkg/ecode"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type CoinValue struct {
	asset  string
	amount decimal.Decimal
}

func NewCoinValue(asset string, amount decimal.Decimal) *CoinValue {
	coinValue := &CoinValue{}
	if err := coinValue.setAsset(asset); err != nil {
		return nil
	}
	if err := coinValue.setAmount(amount); err != nil {
		return nil
	}
	return coinValue
}

func (cv *CoinValue) Asset() string {
	return cv.asset
}

func (cv *CoinValue) setAsset(asset string) error {
	if cv.asset != "" {
		return ecode.DepositAssetChange
	}
	if asset == "" {
		return ecode.DepositAssetInvalid
	}
	cv.asset = asset
	return nil
}

func (cv *CoinValue) Amount() decimal.Decimal {
	return cv.amount
}

func (cv *CoinValue) setAmount(amount decimal.Decimal) error {
	if cv.amount.GreaterThanZero() {
		return ecode.DepositAmountChange
	}
	if !amount.GreaterThanOrEqualZero() {
		return ecode.DepositAmountInvalid
	}
	cv.amount = amount
	return nil
}
