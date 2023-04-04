package deposit

import (
	"github.com/xlalon/golee/pkg/ecode"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type AssetValue struct {
	asset  string
	amount decimal.Decimal
}

func NewAssetValue(asset string, amount decimal.Decimal) *AssetValue {
	coinValue := &AssetValue{}
	if err := coinValue.setAsset(asset); err != nil {
		return nil
	}
	if err := coinValue.setAmount(amount); err != nil {
		return nil
	}
	return coinValue
}

func (cv *AssetValue) Asset() string {
	return cv.asset
}

func (cv *AssetValue) setAsset(asset string) error {
	if cv.asset != "" {
		return ecode.DepositAssetChange
	}
	if asset == "" {
		return ecode.DepositAssetInvalid
	}
	cv.asset = asset
	return nil
}

func (cv *AssetValue) Amount() decimal.Decimal {
	return cv.amount
}

func (cv *AssetValue) setAmount(amount decimal.Decimal) error {
	if cv.amount.GreaterThanZero() {
		return ecode.DepositAmountChange
	}
	if !amount.GreaterThanOrEqualZero() {
		return ecode.DepositAmountInvalid
	}
	cv.amount = amount
	return nil
}
