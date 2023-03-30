package chainasset

import (
	"github.com/xlalon/golee/pkg/ecode"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type Service struct {
	chainRepo ChainRepository
}

func NewService(chainRepo ChainRepository) *Service {
	return &Service{chainRepo: chainRepo}
}

func (s *Service) AddAsset(chainCode, assetCode, assetName, identity string, precession int64) error {
	chain, err := s.chainRepo.GetChainByCode(chainCode)
	if err != nil {
		return err
	}
	if chain == nil {
		return ecode.ChainInvalid
	}
	asset := AssetFactory(&AssetDTO{
		Id:         s.chainRepo.NextId(),
		Code:       assetCode,
		Name:       assetName,
		Chain:      chainCode,
		Identity:   identity,
		Precession: precession,
		Status:     AssetStatusOffline,
		Setting:    nil,
	})
	if asset == nil {
		return ecode.AssetInvalid
	}
	if err = chain.AddAsset(asset); err != nil {
		return err
	}
	if err = s.chainRepo.SaveChain(chain); err != nil {
		return err
	}

	return nil
}

func (s *Service) SetAssetSettings(assetCode, chainCode string, minDepositAmount, withdrawFee, toHotThreshold decimal.Decimal) error {

	asset, err := s.chainRepo.GetAssetByCode(chainCode, assetCode)
	if err != nil {
		return err
	}
	if err = asset.SetSetting(NewAssetSetting(
		minDepositAmount,
		withdrawFee,
		toHotThreshold,
	)); err != nil {
		return err
	}

	return s.chainRepo.SaveAsset(asset)
}
