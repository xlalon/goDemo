package http

import (
	"github.com/xlalon/golee/internal/infra/repository"
	"github.com/xlalon/golee/internal/interface/conf"
	"github.com/xlalon/golee/internal/interface/service"
	"github.com/xlalon/golee/internal/onchain/x"
	"github.com/xlalon/golee/internal/service/chain"
	"github.com/xlalon/golee/internal/service/wallet"
	"github.com/xlalon/golee/pkg/net/http/server"
)

var (
	accountSvc *service.AccountService
	chainSvc   *service.ChainService
	depositSvc *service.DepositService

	repo      *repository.Registry
	walletSvc *wallet.Service
	chainSVC  *chain.Service
)

var (
	accountH = &accountHandler{}
	chainH   = &chainHandler{}
	depositH = &depositHandler{}
)

func Init(r *server.Engine, conf *conf.Config) error {
	initServices(conf)
	registerRouter(r)
	return nil
}

func registerRouter(r *server.Engine) {
	v1 := r.Group("/v1")

	rChain := v1.Group("chain")
	rChain.GET("/:chain/height/latest", chainH.getLatestHeight)
	rChain.GET("/list", chainH.getChains)

	rAsset := v1.Group("asset")
	rAsset.GET("/list", chainH.getAssets)

	rAccount := v1.Group("account")
	rAccount.POST("/new", accountH.newAccount)
	rAccount.GET("/:address/detail", accountH.getAccountDetail)
	rAccount.GET("/:address/balance", accountH.getAccountBalance)
	rAccount.GET("/list", accountH.getAccounts)

	rDeposit := v1.Group("deposit")
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
	accountSvc = service.NewAccountService(repo.WalletRepository())
}
