package chain

import (
	"github.com/xlalon/golee/internal/service/chain/domain"
)

type Chain struct {
	chainSvc *chainasset.Service
}

func NewChain(chainRepo domain.ChainRepository) *Chain {
	return &Chain{
		chainSvc: chainasset.NewService(chainRepo),
	}
}

func (c *Chain) GetHeight(code string) (int64, error) {
	return c.chainSvc.GetChainLatestHeight(code)
}
