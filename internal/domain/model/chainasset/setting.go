package chainasset

import (
	"github.com/xlalon/golee/pkg/ecode"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type AssetSetting struct {
	minDepositAmount decimal.Decimal
	withdrawFee      decimal.Decimal
	toHotThreshold   decimal.Decimal
}

func NewAssetSetting(minDepositAmount, withdrawFee, toHotThreshold decimal.Decimal) *AssetSetting {
	assetSetting := &AssetSetting{}
	if err := assetSetting.setMinDepositAmount(minDepositAmount); err != nil {
		return nil
	}
	if err := assetSetting.setWithdrawFee(withdrawFee); err != nil {
		return nil
	}
	if err := assetSetting.setToHotThreshold(toHotThreshold); err != nil {
		return nil
	}
	return assetSetting
}

func (as *AssetSetting) MinDepositAmount() decimal.Decimal {
	return as.minDepositAmount
}

func (as *AssetSetting) setMinDepositAmount(amount decimal.Decimal) error {
	if as.minDepositAmount.GreaterThanZero() {
		return ecode.ParameterChangeError
	}
	as.minDepositAmount = amount
	return nil
}

func (as *AssetSetting) WithdrawFee() decimal.Decimal {
	return as.withdrawFee
}

func (as *AssetSetting) setWithdrawFee(amount decimal.Decimal) error {
	if as.withdrawFee.GreaterThanZero() {
		return ecode.ParameterChangeError
	}
	as.withdrawFee = amount
	return nil
}

func (as *AssetSetting) ToHotThreshold() decimal.Decimal {
	return as.toHotThreshold
}

func (as *AssetSetting) setToHotThreshold(amount decimal.Decimal) error {
	if as.toHotThreshold.GreaterThanZero() {
		return ecode.ParameterChangeError
	}
	as.toHotThreshold = amount
	return nil
}

func (as *AssetSetting) ToAssetSettingDTO() *AssetSettingDTO {
	return &AssetSettingDTO{
		MinDepositAmount: as.MinDepositAmount(),
		WithdrawFee:      as.WithdrawFee(),
		ToHotThreshold:   as.ToHotThreshold(),
	}
}
