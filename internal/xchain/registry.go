package xchain

import (
	"fmt"
	"sync"

	"github.com/xlalon/golee/internal/xchain/conf"
)

var (
	mu             sync.Mutex
	chainsRegistry = &registry{
		chains:         []Chain{},
		chainToApis:    make(map[Chain]Chainer),
		chainToConfigs: make(map[Chain]*conf.ChainConfig),
	}
)

func RegisterChain(code Chain, chainConfig *conf.ChainConfig, api Chainer) {
	chainsRegistry.addChain(code, chainConfig, api)
}

type registry struct {
	chains         []Chain
	chainToApis    map[Chain]Chainer
	chainToConfigs map[Chain]*conf.ChainConfig
}

func (r *registry) addChain(chain Chain, chainConfig *conf.ChainConfig, api Chainer) {
	chain = chain.Normalize()
	if _, ok := r.chainToApis[chain]; ok {
		panic(fmt.Errorf("chain %s already registered", chain))
	}
	mu.Lock()
	defer mu.Unlock()
	r.chainToConfigs[chain] = chainConfig
	r.chainToApis[chain] = api
	r.chains = append(r.chains, chain)
}

func (r *registry) allChain() []Chain {
	chainsCopy := make([]Chain, len(chainsRegistry.chains))
	chainsCopy = append(chainsCopy, chainsRegistry.chains...)
	return chainsCopy
}

func (r *registry) allChainToApi() map[Chain]Chainer {
	chainToApisCopy := make(map[Chain]Chainer, len(chainsRegistry.chainToApis))
	for chain, api := range chainsRegistry.chainToApis {
		chainToApisCopy[chain] = api
	}
	return chainToApisCopy
}

func (r *registry) getChainApi(chain Chain) (Chainer, bool) {
	api, ok := r.chainToApis[chain.Normalize()]
	return api, ok
}

func (r *registry) getChainConfig(chain Chain) (*conf.ChainConfig, bool) {
	chainConfig, ok := r.chainToConfigs[chain.Normalize()]
	if !ok || chainConfig == nil {
		return nil, ok
	}
	chainConfigCopy := *chainConfig
	return &chainConfigCopy, ok
}
