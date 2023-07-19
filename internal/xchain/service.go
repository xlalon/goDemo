package xchain

import "github.com/xlalon/golee/internal/xchain/conf"

type Service struct {
	registry *registry
}

func NewService() *Service {
	return &Service{registry: chainsRegistry}
}

func (s *Service) Chains() []Chain {
	return s.registry.allChain()
}

func (s *Service) ChainToApis() map[Chain]Chainer {
	return s.registry.allChainToApi()
}

func (s *Service) GetChainApi(chain Chain) (Chainer, bool) {
	return s.registry.getChainApi(chain)
}

func (s *Service) GetChainConfig(chain Chain) (*conf.ChainConfig, bool) {
	return s.registry.getChainConfig(chain)
}
