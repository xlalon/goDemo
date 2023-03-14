package http

import (
	"github.com/xlalon/golee/internal/interface/conf"
	"github.com/xlalon/golee/internal/interface/service"
	"github.com/xlalon/golee/internal/onchain/x"
	"github.com/xlalon/golee/pkg/net/http/server"
)

var (
	accountSvc *service.AccountService
	chainSvc   *service.ChainService
	depositSvc *service.DepositService
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
	chainSvc = service.NewChainService(conf)
	depositSvc = service.NewDepositService(conf)
	accountSvc = service.NewAccountService(conf)
}
