package service

import (
	"errors"
	"gorm.io/gorm"

	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type AssetService struct {
	domainRegistry *Registry
}

func NewAssetService() *AssetService {
	return &AssetService{DomainRegistry}
}

func (s *AssetService) GetAssets() (interface{}, error) {
	return s.assetsToDTOs(s.domainRegistry.chainRepository.GetAssets())
}

func (s *AssetService) GetAssetsByChain(chainCode string) ([]*chainasset.AssetDTO, error) {
	return s.assetsToDTOs(s.domainRegistry.chainRepository.GetChainAssets(chainCode))
}

func (s *AssetService) SetAssetSettings(assetCode, chainCode string, minDepositAmount, withdrawFee, toHotThreshold decimal.Decimal) error {
	return s.domainRegistry.chainAssetSvc.SetAssetSettings(assetCode, chainCode, minDepositAmount, withdrawFee, toHotThreshold)
}

func (s *AssetService) GetAssetSettings(chainCode, assetCode string) (*chainasset.AssetSettingDTO, error) {
	setting, err := s.domainRegistry.chainRepository.GetAssetSetting(chainCode, assetCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			setting = &chainasset.AssetSetting{}
		} else {
			return nil, err
		}
	}
	settingDTO := setting.ToAssetSettingDTO()
	settingDTO.ChainCode = chainCode
	settingDTO.AssetCode = assetCode

	return settingDTO, nil
}

func (s *AssetService) assetToDTO(asset *chainasset.Asset, err error) (*chainasset.AssetDTO, error) {
	if err != nil {
		return nil, err
	}
	return asset.ToAssetDTO(), nil
}

func (s *AssetService) assetsToDTOs(assets []*chainasset.Asset, err error) ([]*chainasset.AssetDTO, error) {
	if err != nil {
		return nil, err
	}
	var assetsDTO []*chainasset.AssetDTO
	for _, asset := range assets {
		assetsDTO = append(assetsDTO, asset.ToAssetDTO())
	}
	return assetsDTO, nil
}