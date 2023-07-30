package chainasset

import (
	"github.com/xlalon/golee/pkg/database/mysql"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type Chain struct {
	mysql.Model

	Code   string
	Name   string
	Status string
}

type Asset struct {
	mysql.Model

	ChainCode string
	Code      string
	Name      string
	Identity  string
	Precision int64
	Status    string
}

type AssetSetting struct {
	mysql.Model

	ChainCode string
	AssetCode string

	MinDepositAmount decimal.Decimal
	WithdrawFee      decimal.Decimal
	ToHotThreshold   decimal.Decimal
}
