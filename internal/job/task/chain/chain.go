package chain

import (
	"github.com/xlalon/golee/internal/service/chain"
	"github.com/xlalon/golee/internal/service/chain/domain"
)

type Chain struct {
	chainSvc *chain.Service
}

func NewChain(chainRepo domain.ChainRepository) *Chain {
	return &Chain{
		chainSvc: chain.NewService(chainRepo),
	}
}

func (c *Chain) GetHeight(code string) (int64, error) {
	return c.chainSvc.GetChainLatestHeight(code)
}
