package service

import (
	"errors"
	"github.com/xlalon/golee/internal/domain"
	"gorm.io/gorm"

	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type AssetService struct {
	Service
}

func NewAssetService() *AssetService {
	return &AssetService{Service{DomainRegistry: domain.DomainRegistry}}
}

func (s *AssetService) GetAssets() (interface{}, error) {
	return s.assetsToDTOs(s.DomainRegistry.ChainRepository.GetAssets())
}

func (s *AssetService) GetAssetsByChain(chainCode string) ([]*chainasset.AssetDTO, error) {
	return s.assetsToDTOs(s.DomainRegistry.ChainRepository.GetChainAssets(chainasset.ChainCode(chainCode)))
}

func (s *AssetService) SetAssetSettings(assetCode, chainCode string, minDepositAmount, withdrawFee, toHotThreshold decimal.Decimal) error {
	cc, ac := chainasset.ChainCode(chainCode), chainasset.AssetCode(assetCode)
	asset, err := s.DomainRegistry.ChainRepository.GetAssetByCode(cc, ac)
	if err != nil {
		return err
	}
	setting, err := asset.ApplySetting(
		minDepositAmount,
		withdrawFee,
		toHotThreshold,
	)
	if err != nil {
		return err
	}

	return s.DomainRegistry.ChainRepository.SaveAssetSetting(cc, ac, setting)
}

func (s *AssetService) GetAssetSettings(chainCode, assetCode string) (*chainasset.AssetSettingDTO, error) {
	cc, ac := chainasset.ChainCode(chainCode), chainasset.AssetCode(assetCode)
	setting, err := s.DomainRegistry.ChainRepository.GetAssetSetting(cc, ac)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			setting = &chainasset.AssetSetting{}
		} else {
			return nil, err
		}
	}
	settingDTO := setting.ToAssetSettingDTO()
	settingDTO.ChainCode = cc
	settingDTO.AssetCode = ac

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
