package model

import "github.com/xlalon/golee/pkg/math/decimal"

type ChainDTO struct {
	Id     int64       `json:"id"`
	Code   string      `json:"code"`
	Name   string      `json:"name"`
	Status string      `json:"status"`
	Assets []*AssetDTO `json:"assets"`
}

type AssetDTO struct {
	Id         int64  `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Chain      string `json:"chain"`
	Identity   string `json:"identity"`
	Precession int64  `json:"precession"`
	Status     string `json:"status"`

	Setting *AssetSettingDTO `json:"-"`
}

type AssetSettingDTO struct {
	ChainCode        string          `json:"chain_code"`
	AssetCode        string          `json:"asset_code"`
	MinDepositAmount decimal.Decimal `json:"min_deposit_amount"`
	WithdrawFee      decimal.Decimal `json:"withdraw_fee"`
	ToHotThreshold   decimal.Decimal `json:"to_hot_threshold"`
}
