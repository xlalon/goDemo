package x

import (
	"sync"

	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/internal/onchain/conf"
	"github.com/xlalon/golee/internal/onchain/x/band"
	"github.com/xlalon/golee/internal/onchain/x/waxp"
)

var (
	once sync.Once
)

func Init(conf *conf.Config) error {

	once.Do(func() {

		bandChain := band.New(conf.Band)
		onchain.RegisterChain(bandChain.Code, bandChain)

		waxChain := waxp.New(conf.Waxp)
		onchain.RegisterChain(waxChain.Code, waxChain)

	})

	return nil
}
