package service

import (
	"github.com/xlalon/golee/internal/domain/chainasset/model"
	"github.com/xlalon/golee/internal/onchain"
)

type ChainService struct {
	domainRegistry *Registry
}

func NewChainService() *ChainService {
	return &ChainService{DomainRegistry}
}

func (s *ChainService) NewChain(code, name string) (interface{}, error) {

	chain := model.ChainFactory(&model.ChainDTO{
		Id:     s.domainRegistry.chainRepository.NextId(),
		Code:   code,
		Name:   name,
		Status: model.ChainStatusOffline,
	}, nil)

	if err := s.domainRegistry.chainRepository.SaveChain(chain); err != nil {
		return nil, err
	}

	chain, err := s.domainRegistry.chainRepository.GetChainByCode(code)
	if err != nil {
		return nil, err
	}
	return chain.ToChainDTO(), nil
}

func (s *AssetService) AddAsset(chainCode string, assetCode, assetName, identity string, precession int64) error {
	return s.domainRegistry.chainAssetService.AddAsset(chainCode, assetCode, assetName, identity, precession)
}

func (s *ChainService) GetChainByCode(chainCode string) (interface{}, error) {
	return s.chainToDTO(s.domainRegistry.chainRepository.GetChainByCode(chainCode))
}

func (s *ChainService) GetChains() (interface{}, error) {
	return s.chainsToDTOs(s.domainRegistry.chainRepository.GetChains())
}

func (s *ChainService) GetNodeInfo(chainCode string) (interface{}, error) {
	chainRpc, _ := s.domainRegistry.onChainService.GetChainApi(onchain.Code(chainCode))
	return chainRpc.GetNodeInfo()
}

func (s *ChainService) GetChainLatestHeight(chainCode string) (int64, error) {
	chainRpc, _ := s.domainRegistry.onChainService.GetChainApi(onchain.Code(chainCode))
	return chainRpc.GetLatestHeight()
}

func (s *ChainService) chainToDTO(chain *model.Chain, err error) (*model.ChainDTO, error) {
	if err != nil {
		return nil, err
	}
	return chain.ToChainDTO(), nil
}

func (s *ChainService) chainsToDTOs(chains []*model.Chain, err error) ([]*model.ChainDTO, error) {
	if err != nil {
		return nil, err
	}
	var chainsDTO []*model.ChainDTO
	for _, chain := range chains {
		chainsDTO = append(chainsDTO, chain.ToChainDTO())
	}
	return chainsDTO, nil
}
