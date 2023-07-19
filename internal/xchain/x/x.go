package x

import (
	"sync"

	"github.com/xlalon/golee/internal/xchain"
	"github.com/xlalon/golee/internal/xchain/conf"
	"github.com/xlalon/golee/internal/xchain/x/band"
	"github.com/xlalon/golee/internal/xchain/x/waxp"
)

var (
	once sync.Once
)

func Init(conf *conf.Config) {

	once.Do(func() {

		bandChain := band.New(conf.Band)
		xchain.RegisterChain(bandChain.Code, bandChain.Config, bandChain)

		waxChain := waxp.New(conf.Waxp)
		xchain.RegisterChain(waxChain.Code, waxChain.Config, waxChain)

	})
}
