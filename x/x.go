package x

import (
	"fmt"
	"sync"

	"github.com/xlalon/golee/x/chain"
	"github.com/xlalon/golee/x/chain/band"
	"github.com/xlalon/golee/x/conf"
)

var (
	once     sync.Once
	mu       sync.Mutex
	registry = make(map[string]chain.Chainer)
)

func Init(conf *conf.ChainConfig) {
	once.Do(func() {
		bandChain := band.NewBand(conf.Band)
		register(bandChain.Code, bandChain)
	})
}

func register(code string, chain chain.Chainer) {
	if _, ok := registry[code]; ok {
		panic(fmt.Sprintf("Duplicate Chain, %s", code))
	}
	mu.Lock()
	defer mu.Unlock()
	registry[code] = chain
}

func GetChain(code string) (chain.Chainer, error) {
	c, ok := registry[code]
	if !ok {
		return nil, fmt.Errorf(fmt.Sprintf("Chain not found, %s", code))
	}
	return c, nil
}
