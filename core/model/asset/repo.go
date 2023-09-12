package asset

import "github.com/xlalon/golee/core/model/chain"

type Repo interface {
	Save(a *Asset) error

	GetAssetById(id int64) (*Asset, error)
	GetAssetByCode(code Code) (*Asset, error)
	GetAssets() ([]*Asset, error)
	GetAssetsByChain(chainCode chain.Code) ([]*Asset, error)
}
