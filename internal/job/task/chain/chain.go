package chain

import (
	"context"
	"github.com/xlalon/golee/internal/domain"
	"github.com/xlalon/golee/internal/xchain"
	"github.com/xlalon/golee/pkg/ecode"
)

type Chain struct {
	DomainRegistry *domain.Registry
}

func NewChain() *Chain {
	return &Chain{
		DomainRegistry: domain.DomainRegistry,
	}
}

func (c *Chain) GetHeight(code string) (int64, error) {
	cApi, ok := c.DomainRegistry.OnChainSvc.GetChainApi(xchain.Chain(code))
	if !ok {
		return 0, ecode.ChainNotFound
	}
	return cApi.GetLatestHeight(context.Background())
}
