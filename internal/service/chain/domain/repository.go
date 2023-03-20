package domain

import "github.com/xlalon/golee/pkg/database/mysql"

type ChainRepository interface {
	mysql.IdGeneratorRepository

	Save(*Chain) error
	GetChains() ([]*Chain, error)
	GetChainByCode(chainCode string) (*Chain, error)
	GetChainByAsset(assetCode string) (*Chain, error)

	GetAssets() ([]*Asset, error)
	GetAssetByCode(assetCode string) (*Asset, error)
	GetAssetByIdentity(chainCode, identity string) (*Asset, error)
	GetAssetsByChain(chainCode string) ([]*Asset, error)
}
