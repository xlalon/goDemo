package service

import (
	"errors"

	"gorm.io/gorm"

	"github.com/xlalon/golee/internal/domain/chainasset/model"
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

func (s *AssetService) GetAssetsByChain(chainCode string) ([]*model.AssetDTO, error) {
	return s.assetsToDTOs(s.domainRegistry.chainRepository.GetChainAssets(chainCode))
}

func (s *AssetService) SetAssetSettings(assetCode, chainCode string, minDepositAmount, withdrawFee, toHotThreshold decimal.Decimal) error {
	return s.domainRegistry.chainAssetService.SetAssetSettings(assetCode, chainCode, minDepositAmount, withdrawFee, toHotThreshold)
}

func (s *AssetService) GetAssetSettings(chainCode, assetCode string) (*model.AssetSettingDTO, error) {
	setting, err := s.domainRegistry.chainRepository.GetAssetSetting(chainCode, assetCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			setting = &model.AssetSetting{}
		} else {
			return nil, err
		}
	}
	settingDTO := setting.ToAssetSettingDTO()
	settingDTO.ChainCode = chainCode
	settingDTO.AssetCode = assetCode

	return settingDTO, nil
}

func (s *AssetService) assetToDTO(asset *model.Asset, err error) (*model.AssetDTO, error) {
	if err != nil {
		return nil, err
	}
	return asset.ToAssetDTO(), nil
}

func (s *AssetService) assetsToDTOs(assets []*model.Asset, err error) ([]*model.AssetDTO, error) {
	if err != nil {
		return nil, err
	}
	var assetsDTO []*model.AssetDTO
	for _, asset := range assets {
		assetsDTO = append(assetsDTO, asset.ToAssetDTO())
	}
	return assetsDTO, nil
}
