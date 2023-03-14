package chain

import (
	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/internal/service/chain/conf"
	"github.com/xlalon/golee/internal/service/chain/domain"
	"github.com/xlalon/golee/internal/service/chain/repository"
	"github.com/xlalon/golee/internal/service/chain/repository/dao"
)

type Service struct {
	repo       repository.ChainRepository
	onchainSvc *onchain.Service
}

func NewService(conf *conf.Config) *Service {
	return &Service{
		repo:       dao.New(conf),
		onchainSvc: onchain.NewService(),
	}
}

func (s *Service) GetChains() ([]*domain.ChainDTO, error) {
	return s.chainsDMToDTO(s.repo.GetChains())
}

func (s *Service) GetChainByCode(chainCode string) (*domain.ChainDTO, error) {
	return s.chainDMToDTO(s.repo.GetChainByCode(chainCode))
}

func (s *Service) GetAssetsByChain(chainCode string) ([]*domain.AssetDTO, error) {
	return s.assetsDMToDTOs(s.repo.GetAssetsByChain(chainCode))
}

func (s *Service) GetAssets() ([]*domain.AssetDTO, error) {
	return s.assetsDMToDTOs(s.repo.GetAssets())
}

func (s *Service) GetAssetByIdentity(chainCode, identity string) (*domain.AssetDTO, error) {
	return s.assetDMToDTO(s.repo.GetAssetByIdentity(chainCode, identity))
}

func (s *Service) GetChainByAsset(assetCode string) (*domain.ChainDTO, error) {
	return s.chainDMToDTO(s.repo.GetChainByAsset(assetCode))
}

func (s *Service) GetChainLatestHeight(chainCode string) (int64, error) {
	chainRpc, _ := s.onchainSvc.GetChainApi(onchain.Code(chainCode))
	return chainRpc.GetLatestHeight()
}

func (s *Service) chainDMToDTO(chainDM *domain.Chain, err error) (*domain.ChainDTO, error) {
	if err != nil {
		return nil, err
	}
	return chainDM.ToChainDTO(), nil
}

func (s *Service) chainsDMToDTO(assetsDM []*domain.Chain, err error) ([]*domain.ChainDTO, error) {
	if err != nil {
		return nil, err
	}
	var chainsDTO []*domain.ChainDTO
	for _, assetDM := range assetsDM {
		chainsDTO = append(chainsDTO, assetDM.ToChainDTO())
	}
	return chainsDTO, nil
}

func (s *Service) assetDMToDTO(assetDM *domain.Asset, err error) (*domain.AssetDTO, error) {
	if err != nil {
		return nil, err
	}
	return assetDM.ToAssetDTO(), nil
}

func (s *Service) assetsDMToDTOs(assetsDM []*domain.Asset, err error) ([]*domain.AssetDTO, error) {
	if err != nil {
		return nil, err
	}
	var assetsDTO []*domain.AssetDTO
	for _, assetDM := range assetsDM {
		assetsDTO = append(assetsDTO, assetDM.ToAssetDTO())
	}
	return assetsDTO, nil
}
