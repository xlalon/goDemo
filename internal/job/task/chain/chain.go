package chain

import (
	"github.com/xlalon/golee/internal/job/conf"
	"github.com/xlalon/golee/internal/service/chain"
	chainConf "github.com/xlalon/golee/internal/service/chain/conf"
)

type Chain struct {
	chainSvc *chain.Service
}

func NewChain(conf *conf.Config) *Chain {
	return &Chain{
		chainSvc: chain.NewService(&chainConf.Config{
			Mysql: conf.Mysql,
			Redis: conf.Redis,
		}),
	}
}

func (c *Chain) GetHeight(code string) (int64, error) {
	return c.chainSvc.GetChainLatestHeight(code)
}
