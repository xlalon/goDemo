package onchain

import (
	"fmt"

	"github.com/xlalon/golee/internal/onchain/conf"
)

var (
	_defaultChainRegistry = &registry{
		chains:         []Code{},
		chainToApis:    make(map[Code]Chainer),
		chainToConfigs: make(map[Code]*conf.ChainConfig),
	}
)

func RegisterChain(code Code, chainConfig *conf.ChainConfig, api Chainer) {
	_defaultChainRegistry.addChain(code, chainConfig, api)
}

type registry struct {
	chains         []Code
	chainToApis    map[Code]Chainer
	chainToConfigs map[Code]*conf.ChainConfig
}

func (r *registry) addChain(code Code, chainConfig *conf.ChainConfig, api Chainer) {
	if _, ok := r.chainToApis[code]; ok {
		panic(fmt.Errorf("chain %s already registered", code))
	}
	r.chainToConfigs[code] = chainConfig
	r.chainToApis[code] = api
	r.chains = append(r.chains, code)
}

func (r *registry) allChain() []Code {
	var chainsCopy []Code
	chainsCopy = append(chainsCopy, _defaultChainRegistry.chains...)
	return chainsCopy
}

func (r *registry) allChainToApi() map[Code]Chainer {
	chainToApisCopy := make(map[Code]Chainer, len(_defaultChainRegistry.chainToApis))
	for code, api := range _defaultChainRegistry.chainToApis {
		chainToApisCopy[code] = api
	}
	return chainToApisCopy
}

func (r *registry) getChainApi(code Code) (Chainer, bool) {
	api, ok := r.chainToApis[code]
	return api, ok
}

func (r *registry) getChainConfig(code Code) (*conf.ChainConfig, bool) {
	chainConfig, ok := r.chainToConfigs[code]
	return chainConfig, ok
}
