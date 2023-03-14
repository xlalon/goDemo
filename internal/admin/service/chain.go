package service

import (
	"github.com/xlalon/golee/internal/admin/conf"
	"github.com/xlalon/golee/internal/service/chain"
	chainConf "github.com/xlalon/golee/internal/service/chain/conf"
)

type ChainService struct {
	assetSvc *chain.Service
}

func NewChainService(conf *conf.Config) *ChainService {
	return &ChainService{
		assetSvc: chain.NewService(&chainConf.Config{
			Mysql: conf.Mysql,
			Redis: conf.Redis,
		}),
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
