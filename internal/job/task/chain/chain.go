package chain

import (
	"context"
	"github.com/xlalon/golee/internal/domain"
	"github.com/xlalon/golee/internal/onchain"
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
	cApi, ok := c.DomainRegistry.OnChainSvc.GetChainApi(onchain.Code(code))
	if !ok {
		return 0, ecode.ChainNotFound
	}
	return cApi.GetLatestHeight(context.Background())
}
