package chain

import (
	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/pkg/ecode"
)

type Chain struct {
	onchainSvc *onchain.Service
}

func NewChain() *Chain {
	return &Chain{
		onchainSvc: onchain.NewService(),
	}
}

func (c *Chain) GetHeight(code string) (int64, error) {
	cApi, ok := c.onchainSvc.GetChainApi(onchain.Code(code))
	if !ok {
		return 0, ecode.ChainNotFound
	}
	return cApi.GetLatestHeight()
}
