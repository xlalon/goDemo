package chain

import (
	"fmt"

	"github.com/xlalon/golee/internal/job/task/chain"
	"github.com/xlalon/golee/internal/xchain"
	"github.com/xlalon/golee/pkg/job/worker"
)

type Chain struct {
	chainTask  *chain.Chain
	onchainSvc *xchain.Service
}

func NewChain() *Chain {
	return &Chain{
		chainTask:  chain.NewChain(),
		onchainSvc: xchain.NewService(),
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
