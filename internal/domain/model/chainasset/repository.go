package chainasset

import (
	"github.com/xlalon/golee/pkg/database/mysql"
)

type ChainRepository interface {
	mysql.IdGeneratorRepository

	SaveChain(*Chain) error
	GetChains() ([]*Chain, error)
	GetChainByCode(chainCode string) (*Chain, error)
	GetAssetChains(assetCode string) ([]*Chain, error)

	SaveAsset(asset *Asset) error
	GetAssets() ([]*Asset, error)
	GetAssetsByCode(assetCode string) ([]*Asset, error)
	GetAssetByCode(chainCode, assetCode string) (*Asset, error)
	GetAssetByIdentity(chainCode, identity string) (*Asset, error)
	GetChainAssets(chainCode string) ([]*Asset, error)

	SaveAssetSetting(chainCode, assetCode string, settings *AssetSetting) error
	GetAssetSettings() ([]*AssetSetting, error)
	GetAssetSettingsByChain(chainCode string) ([]*AssetSetting, error)
	GetAssetSetting(chainCode, assetCode string) (*AssetSetting, error)
}
