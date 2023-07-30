package job

import (
	"github.com/xlalon/golee/internal/domain"
	"github.com/xlalon/golee/internal/infra/repository"
	raccount "github.com/xlalon/golee/internal/infra/repository/account"
	rchainasset "github.com/xlalon/golee/internal/infra/repository/chainasset"
	rdeposit "github.com/xlalon/golee/internal/infra/repository/deposit"
	"github.com/xlalon/golee/internal/job/conf"
	"github.com/xlalon/golee/internal/job/scheduler/chain"
	"github.com/xlalon/golee/internal/xchain/x"
	"github.com/xlalon/golee/pkg/job/worker"
)

var (
	chainScd *chain.Chain
)

func Init(server *worker.Server, conf *conf.Config) error {
	x.Init(conf.Chain)
	initDomain(conf)
	initSchedules()
	return registerTask(server)
}

func initDomain(conf *conf.Config) {
	_registry := repository.NewRegistry(&repository.Config{
		Chain: &rchainasset.Config{
			Mysql: conf.Mysql,
			Redis: conf.Redis,
		},
		Deposit: &rdeposit.Config{
			Mysql: conf.Mysql,
			Redis: conf.Redis,
		},
		Account: &raccount.Config{
			Mysql: conf.Mysql,
			Redis: conf.Redis,
		},
	})
	domain.Init(_registry.ChainRepository(), _registry.DepositRepository(), _registry.AccountRepository())
}

func initSchedules() {
	chainScd = chain.NewChain()
}

func registerTask(server *worker.Server) error {
	return chainScd.Register(server)
}
