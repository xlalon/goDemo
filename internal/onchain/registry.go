package onchain

import "fmt"

var (
	defaultChainRegistry = &registry{
		chains:      []Code{},
		chainToApis: make(map[Code]Chainer),
	}
)

type registry struct {
	chains      []Code
	chainToApis map[Code]Chainer
}

func (r *registry) addChain(c Code, api Chainer) {
	if _, ok := r.chainToApis[c]; ok {
		panic(fmt.Errorf("chain %s already registered", c))
	}
	r.chainToApis[c] = api
	r.chains = append(r.chains, c)
}

func (r *registry) allChain() []Code {
	var chainsCopy []Code
	chainsCopy = append(chainsCopy, defaultChainRegistry.chains...)
	return chainsCopy
}

func (r *registry) allChainToApi() map[Code]Chainer {
	chainToApisCopy := make(map[Code]Chainer, len(defaultChainRegistry.chainToApis))
	for code, api := range defaultChainRegistry.chainToApis {
		chainToApisCopy[code] = api
	}
	return chainToApisCopy
}

func (r *registry) getChainApi(c Code) (Chainer, bool) {
	api, ok := r.chainToApis[c]
	return api, ok
}

func RegisterChain(c Code, api Chainer) {
	defaultChainRegistry.addChain(c, api)
}
