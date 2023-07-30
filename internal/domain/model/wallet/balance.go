package wallet

import (
	"fmt"
	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/pkg/ecode"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type AssetValue struct {
	asset  chainasset.AssetCode
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

func (cv *AssetValue) Asset() chainasset.AssetCode {
	return cv.asset
}

func (cv *AssetValue) setAsset(asset string) error {
	if cv.asset != "" {
		return ecode.DepositAssetChange
	}
	if asset == "" {
		return ecode.DepositAssetInvalid
	}
	cv.asset = chainasset.AssetCode(asset)
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

func (cv *AssetValue) String() string {
	return fmt.Sprintf("%v %s", cv.Amount(), cv.Asset())
}

func (cv *AssetValue) ToAssetValueDTO() *BalanceDTO {
	return &BalanceDTO{
		Amount: cv.Amount(),
		Asset:  string(cv.Asset()),
	}
}
