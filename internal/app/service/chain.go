package service

import (
	"context"
	"github.com/xlalon/golee/internal/domain"
	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/pkg/ecode"
)

type ChainService struct {
	Service
}

func NewChainService() *ChainService {
	return &ChainService{Service{DomainRegistry: domain.DomainRegistry}}
}

func (s *ChainService) NewChain(code, name string) (interface{}, error) {

	chain := chainasset.ChainFactory(&chainasset.ChainDTO{
		Id:     s.DomainRegistry.ChainRepository.NextId(),
		Code:   chainasset.ChainCode(code),
		Name:   name,
		Status: chainasset.ChainStatusOffline,
	}, nil)

	if err := s.DomainRegistry.ChainRepository.SaveChain(chain); err != nil {
		return nil, err
	}

	chain, err := s.DomainRegistry.ChainRepository.GetChainByCode(chainasset.ChainCode(code))
	if err != nil {
		return nil, err
	}
	return chain.ToChainDTO(), nil
}

func (s *AssetService) RegisterAsset(chainCode string, assetCode, assetName, identity string, precession int64) error {
	chain, err := s.DomainRegistry.ChainRepository.GetChainByCode(chainasset.ChainCode(chainCode))
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
	return s.DomainRegistry.ChainRepository.SaveAsset(asset)
}

func (s *ChainService) GetChainByCode(chainCode string) (interface{}, error) {
	return s.chainToDTO(s.DomainRegistry.ChainRepository.GetChainByCode(chainasset.ChainCode(chainCode)))
}

func (s *ChainService) GetChains() (interface{}, error) {
	return s.chainsToDTOs(s.DomainRegistry.ChainRepository.GetChains())
}

func (s *ChainService) GetNodeInfo(ctx context.Context, chainCode string) (interface{}, error) {
	chainRpc, _ := s.DomainRegistry.OnChainSvc.GetChainApi(onchain.Code(chainCode))
	return chainRpc.GetNodeInfo(ctx)
}

func (s *ChainService) GetChainLatestHeight(ctx context.Context, chainCode string) (int64, error) {
	chainRpc, _ := s.DomainRegistry.OnChainSvc.GetChainApi(onchain.Code(chainCode))
	return chainRpc.GetLatestHeight(ctx)
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
