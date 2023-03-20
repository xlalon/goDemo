package service

import (
	"github.com/xlalon/golee/internal/service/chain"
	"github.com/xlalon/golee/internal/service/chain/domain"
)

type ChainService struct {
	assetSvc *chain.Service
}

func NewChainService(chainRepo domain.ChainRepository) *ChainService {
	return &ChainService{
		assetSvc: chain.NewService(chainRepo),
	}
}

func (s *ChainService) GetLatestHeight(chainCode string) int64 {
	height, err := s.assetSvc.GetChainLatestHeight(chainCode)
	if err != nil {
		return 0
	}
	return height
}

func (s *ChainService) GetChains() (interface{}, error) {
	return s.assetSvc.GetChains()
}

func (s *ChainService) GetAssets() (interface{}, error) {
	return s.assetSvc.GetAssets()
}
