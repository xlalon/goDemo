package chain

import (
	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/internal/service/chain/domain"
)

type Service struct {
	chainRepo  domain.ChainRepository
	onchainSvc *onchain.Service
}

func NewService(chainRepo domain.ChainRepository) *Service {
	return &Service{
		chainRepo:  chainRepo,
		onchainSvc: onchain.NewService(),
	}
}

func (s *Service) NewChain(chainDTO *domain.ChainDTO) {
	var assets []*domain.Asset
	for _, assetDTO := range chainDTO.Assets {
		assets = append(assets, domain.AssetFactory(assetDTO))
	}
	s.chainRepo.Save(domain.ChainFactory(chainDTO, assets))
}

func (s *Service) AddAsset(chainCode string, assetDTO *domain.AssetDTO) (*domain.ChainDTO, error) {
	chain, err := s.chainRepo.GetChainByCode(chainCode)
	if err != nil {
		return nil, err
	}
	chain.AddAsset(domain.AssetFactory(assetDTO))
	s.chainRepo.Save(chain)
	return s.chainDMToDTO(s.chainRepo.GetChainByCode(chainCode))
}

func (s *Service) GetChains() ([]*domain.ChainDTO, error) {
	return s.chainsDMToDTO(s.chainRepo.GetChains())
}

func (s *Service) GetChainByCode(chainCode string) (*domain.ChainDTO, error) {
	return s.chainDMToDTO(s.chainRepo.GetChainByCode(chainCode))
}

func (s *Service) GetAssetsByChain(chainCode string) ([]*domain.AssetDTO, error) {
	return s.assetsDMToDTOs(s.chainRepo.GetAssetsByChain(chainCode))
}

func (s *Service) GetAssets() ([]*domain.AssetDTO, error) {
	return s.assetsDMToDTOs(s.chainRepo.GetAssets())
}

func (s *Service) GetAssetByIdentity(chainCode, identity string) (*domain.AssetDTO, error) {
	return s.assetDMToDTO(s.chainRepo.GetAssetByIdentity(chainCode, identity))
}

func (s *Service) GetChainByAsset(assetCode string) (*domain.ChainDTO, error) {
	return s.chainDMToDTO(s.chainRepo.GetChainByAsset(assetCode))
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
