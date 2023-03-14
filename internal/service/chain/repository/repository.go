package repository

import (
	"github.com/xlalon/golee/internal/service/chain/domain"
)

type ChainRepository interface {
	NewChain(c *domain.ChainDTO) error
	GetChains() ([]*domain.Chain, error)
	GetChainByCode(chainCode string) (*domain.Chain, error)
	GetChainByAsset(assetCode string) (*domain.Chain, error)

	NewAsset(a *domain.AssetDTO) error
	GetAssets() ([]*domain.Asset, error)
	GetAssetByCode(assetCode string) (*domain.Asset, error)
	GetAssetByIdentity(chainCode, identity string) (*domain.Asset, error)
	GetAssetsByChain(chainCode string) ([]*domain.Asset, error)
}
