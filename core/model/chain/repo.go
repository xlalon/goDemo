package chain

import "github.com/xlalon/golee/core/model/asset"

type Repo interface {
	Save(c *Chain) error

	GetChainById(id int64) (*Chain, error)
	GetChainByCode(code Code) (*Chain, error)
	GetChains() ([]*Chain, error)
	GetAssetsByChainId(id int64) ([]*asset.Asset, error)
	GetAssetsByChainCode(code Code) ([]*asset.Asset, error)
}
