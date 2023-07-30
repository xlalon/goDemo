package chainasset

import (
	"github.com/xlalon/golee/pkg/database/mysql"
)

type ChainRepository interface {
	mysql.IdGeneratorRepository

	SaveChain(*Chain) error
	GetChains() ([]*Chain, error)
	GetChainByCode(chainCode ChainCode) (*Chain, error)
	GetAssetChains(assetCode AssetCode) ([]*Chain, error)

	SaveAsset(asset *Asset) error
	GetAssets() ([]*Asset, error)
	GetAssetsByCode(assetCode AssetCode) ([]*Asset, error)
	GetAssetByCode(chainCode ChainCode, assetCode AssetCode) (*Asset, error)
	GetAssetByIdentity(chainCode ChainCode, identity string) (*Asset, error)
	GetChainAssets(chainCode ChainCode) ([]*Asset, error)

	SaveAssetSetting(chainCode ChainCode, assetCode AssetCode, settings *AssetSetting) error
	GetAssetSettings() ([]*AssetSetting, error)
	GetAssetSettingsByChain(chainCode ChainCode) ([]*AssetSetting, error)
	GetAssetSetting(chainCode ChainCode, assetCode AssetCode) (*AssetSetting, error)
}
