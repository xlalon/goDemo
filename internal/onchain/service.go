package onchain

import "github.com/xlalon/golee/internal/onchain/conf"

type Service struct {
	chainRegistry *registry
}

func NewService() *Service {
	return &Service{chainRegistry: _defaultChainRegistry}
}

func (s *Service) Chains() []Code {
	return s.chainRegistry.allChain()
}

func (s *Service) ChainToApis() map[Code]Chainer {
	return s.chainRegistry.allChainToApi()
}

func (s *Service) GetChainApi(code Code) (Chainer, bool) {
	return s.chainRegistry.getChainApi(code)
}

func (s *Service) GetChainConfig(code Code) (*conf.ChainConfig, bool) {
	return s.chainRegistry.getChainConfig(code)
}
