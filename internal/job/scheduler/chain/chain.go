package chain

import (
	"fmt"

	"github.com/xlalon/golee/internal/job/conf"
	"github.com/xlalon/golee/internal/job/task/chain"
	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/pkg/job/worker"
)

type Chain struct {
	chainTask  *chain.Chain
	onchainSvc *onchain.Service
}

func NewChain(conf *conf.Config) *Chain {
	return &Chain{
		chainTask:  chain.NewChain(conf),
		onchainSvc: onchain.NewService(),
	}
}

func (c *Chain) Register(server *worker.Server) error {
	return c.registerGetHeight(server)
}

func (c *Chain) registerGetHeight(server *worker.Server) error {
	if err := server.RegisterTask("getHeight", c.chainTask.GetHeight); err != nil {
		return err
	}

	for _, cCode := range c.onchainSvc.Chains() {
		signature := &worker.Signature{
			Name: "getHeight",
			Args: []worker.Arg{
				{
					Type:  "string",
					Value: cCode,
				},
			},
		}
		tName := fmt.Sprintf("xchain-getHeight-%s", cCode)
		if err := server.RegisterPeriodicTask("@every 20s", tName, signature); err != nil {
			return err
		}
	}

	return nil
}