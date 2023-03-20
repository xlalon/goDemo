package job

import (
	"github.com/xlalon/golee/internal/infra/repository"
	"github.com/xlalon/golee/internal/job/conf"
	"github.com/xlalon/golee/internal/job/scheduler/chain"
	"github.com/xlalon/golee/internal/onchain/x"
	"github.com/xlalon/golee/pkg/job/worker"
)

var (
	chainScd *chain.Chain

	repo *repository.Registry
)

func Init(server *worker.Server, conf *conf.Config) error {
	initScd(conf)
	return registerTask(server)
}

func initScd(conf *conf.Config) {
	_ = x.Init(conf.Chain)
	repo = repository.NewRegistry(conf.Repository)
	chainScd = chain.NewChain(repo)
}

func registerTask(server *worker.Server) error {
	return chainScd.Register(server)
}
