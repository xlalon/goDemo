package job

import (
	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/internal/domain/model/deposit"
	"github.com/xlalon/golee/internal/domain/model/wallet"
	chainRepo "github.com/xlalon/golee/internal/infra/repository/chainasset"
	depositRepo "github.com/xlalon/golee/internal/infra/repository/deposit"
	walletRepo "github.com/xlalon/golee/internal/infra/repository/wallet"
	"github.com/xlalon/golee/internal/job/conf"
	"github.com/xlalon/golee/internal/job/scheduler/chain"
	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/internal/onchain/x"
	"github.com/xlalon/golee/pkg/job/worker"
)

var (
	chainScd *chain.Chain
)

var DomainRegistry = &Registry{}

type Registry struct {
	chainRepository   chainasset.ChainRepository
	depositRepository deposit.DepositRepository
	walletRepository  wallet.WalletRepository

	onChainSvc *onchain.Service
}

func Init(server *worker.Server, conf *conf.Config) error {
	initScd(conf)
	return registerTask(server)
}

func initScd(conf *conf.Config) {
	x.Init(conf.Chain)

	chainRepository := chainRepo.NewDao(&chainRepo.Config{
		Mysql: conf.Mysql,
		Redis: conf.Redis,
	})

	depositRepository := depositRepo.NewDao(&depositRepo.Config{
		Mysql: conf.Mysql,
		Redis: conf.Redis,
	})

	walletRepository := walletRepo.NewDao(&walletRepo.Config{
		Mysql: conf.Mysql,
		Redis: conf.Redis,
	})

	DomainRegistry = &Registry{
		chainRepository:   chainRepository,
		depositRepository: depositRepository,
		walletRepository:  walletRepository,

		onChainSvc: onchain.NewService(),
	}

	chainScd = chain.NewChain()
}

func registerTask(server *worker.Server) error {
	return chainScd.Register(server)
}
