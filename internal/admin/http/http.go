package http

import (
	"github.com/xlalon/golee/internal/admin/conf"
	"github.com/xlalon/golee/internal/admin/service"
	"github.com/xlalon/golee/internal/infra/repository"
	"github.com/xlalon/golee/internal/onchain/x"
	"github.com/xlalon/golee/internal/service/chain"
	"github.com/xlalon/golee/internal/service/wallet"
	"github.com/xlalon/golee/pkg/net/http/server"
)

var (
	chainSvc   *service.ChainService
	depositSvc *service.DepositService

	repo      *repository.Registry
	walletSvc *wallet.Service
	chainSVC  *chain.Service
)

var (
	chainH   = &chainHandler{}
	depositH = &depositHandler{}
)

func Init(r *server.Engine, conf *conf.Config) error {
	initServices(conf)
	registerRouter(r)
	return nil
}

func registerRouter(r *server.Engine) {
	v1 := r.Group("/v1/admin")

	rChain := v1.Group("/chain")
	rChain.GET("/:chain/height/latest", chainH.getLatestHeight)

	rDeposit := v1.Group("/deposit")
	rDeposit.GET("/:id", depositH.getDeposit)
	rDeposit.GET("/list", depositH.getDeposits)
}

func initServices(conf *conf.Config) {
	_ = x.Init(conf.Chain)
	repo = repository.NewRegistry(conf.Repository)
	chainSVC = chain.NewService(repo.ChainRepository())
	walletSvc = wallet.NewService(repo.WalletRepository())
	chainSvc = service.NewChainService(repo.ChainRepository())
	depositSvc = service.NewDepositService(repo.DepositRepository(), chainSVC, walletSvc)
}
