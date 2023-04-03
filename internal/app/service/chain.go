package service

import (
	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/pkg/ecode"
)

type ChainService struct {
	domainRegistry *Registry
}

func NewChainService() *ChainService {
	return &ChainService{DomainRegistry}
}

func (s *ChainService) NewChain(code, name string) (interface{}, error) {

	chain := chainasset.ChainFactory(&chainasset.ChainDTO{
		Id:     s.domainRegistry.chainRepository.NextId(),
		Code:   chainasset.ChainCode(code),
		Name:   name,
		Status: chainasset.ChainStatusOffline,
	}, nil)

	if err := s.domainRegistry.chainRepository.SaveChain(chain); err != nil {
		return nil, err
	}

	chain, err := s.domainRegistry.chainRepository.GetChainByCode(chainasset.ChainCode(code))
	if err != nil {
		return nil, err
	}
	return chain.ToChainDTO(), nil
}

func (s *AssetService) RegisterAsset(chainCode string, assetCode, assetName, identity string, precession int64) error {
	chain, err := s.domainRegistry.chainRepository.GetChainByCode(chainasset.ChainCode(chainCode))
	if err != nil {
		return err
	}
	if chain == nil {
		return ecode.ChainInvalid
	}
	asset, err := chain.RegisterAsset(chainasset.AssetCode(assetCode), assetName, identity, precession)
	if err != nil {
		return err
	}
	if asset == nil {
		return ecode.AssetInvalid
	}
	return s.domainRegistry.chainRepository.SaveAsset(asset)
}

func (s *ChainService) GetChainByCode(chainCode string) (interface{}, error) {
	return s.chainToDTO(s.domainRegistry.chainRepository.GetChainByCode(chainasset.ChainCode(chainCode)))
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

func (s *ChainService) chainToDTO(chain *chainasset.Chain, err error) (*chainasset.ChainDTO, error) {
	if err != nil {
		return nil, err
	}
	return chain.ToChainDTO(), nil
}

func (s *ChainService) chainsToDTOs(chains []*chainasset.Chain, err error) ([]*chainasset.ChainDTO, error) {
	if err != nil {
		return nil, err
	}
	var chainsDTO []*chainasset.ChainDTO
	for _, chain := range chains {
		chainsDTO = append(chainsDTO, chain.ToChainDTO())
	}
	return chainsDTO, nil
}
