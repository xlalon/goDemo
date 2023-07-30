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
	setting := &AssetSetting{}
	if err := setting.setMinDepositAmount(minDepositAmount); err != nil {
		return nil
	}
	if err := setting.setWithdrawFee(withdrawFee); err != nil {
		return nil
	}
	if err := setting.setToHotThreshold(toHotThreshold); err != nil {
		return nil
	}
	return setting
}

func (as *AssetSetting) MinDepositAmount() decimal.Decimal {
	return as.minDepositAmount
}

func (as *AssetSetting) setMinDepositAmount(amount decimal.Decimal) error {
	if as.minDepositAmount.GreaterThanZero() {
		return ecode.AssetSettingChange
	}
	as.minDepositAmount = amount
	return nil
}

func (as *AssetSetting) WithdrawFee() decimal.Decimal {
	return as.withdrawFee
}

func (as *AssetSetting) setWithdrawFee(amount decimal.Decimal) error {
	if as.withdrawFee.GreaterThanZero() {
		return ecode.AssetSettingChange
	}
	as.withdrawFee = amount
	return nil
}

func (as *AssetSetting) ToHotThreshold() decimal.Decimal {
	return as.toHotThreshold
}

func (as *AssetSetting) setToHotThreshold(amount decimal.Decimal) error {
	if as.toHotThreshold.GreaterThanZero() {
		return ecode.AssetSettingChange
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

type AssetSettingDTO struct {
	ChainCode        ChainCode       `json:"chain_code"`
	AssetCode        AssetCode       `json:"asset_code"`
	MinDepositAmount decimal.Decimal `json:"min_deposit_amount"`
	WithdrawFee      decimal.Decimal `json:"withdraw_fee"`
	ToHotThreshold   decimal.Decimal `json:"to_hot_threshold"`
}
