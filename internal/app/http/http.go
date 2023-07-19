package http

import (
	"github.com/xlalon/golee/internal/app/conf"
	"github.com/xlalon/golee/internal/app/service"
	"github.com/xlalon/golee/internal/domain"
	"github.com/xlalon/golee/internal/infra/repository"
	raccount "github.com/xlalon/golee/internal/infra/repository/account"
	rchainasset "github.com/xlalon/golee/internal/infra/repository/chainasset"
	rdeposit "github.com/xlalon/golee/internal/infra/repository/deposit"
	"github.com/xlalon/golee/internal/xchain/x"
	"github.com/xlalon/golee/pkg/net/http/server"
)

var (
	_assetHandler   = &assetHandler{}
	_chainHandler   = &chainHandler{}
	_depositHandler = &depositHandler{}
	_accountHandler = &accountHandler{}
)

var (
	assetSvc   *service.AssetService
	chainSvc   *service.ChainService
	depositSvc *service.DepositService
	walletSvc  *service.WalletService
)

func Init(r *server.Engine, conf *conf.Config) error {
	x.Init(conf.Chain)
	initDomain(conf)
	initServices()
	registerRouter(r)
	return nil
}

func registerRouter(r *server.Engine) {
	v1 := r.Group("/v1")

	gAsset := v1.Group("asset")
	gAsset.GET("/list", _assetHandler.getAssets)

	gChain := v1.Group("chain")
	gChain.GET("/:chain/height/latest", _chainHandler.getLatestHeight)
	gChain.GET("/list", _chainHandler.getChains)

	gDeposit := v1.Group("deposit")
	gDeposit.GET("/:id", _depositHandler.getDeposit)
	gDeposit.GET("/list", _depositHandler.getDeposits)

	gAccount := v1.Group("account")
	gAccount.POST("/new", _accountHandler.newAccount)
	gAccount.GET("/:address/balance", _accountHandler.getAccountBalance)
	gAccount.GET("/list", _accountHandler.getAccounts)

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

func initServices() {
	chainSvc = service.NewChainService()
	assetSvc = service.NewAssetService()
	depositSvc = service.NewDepositService()
	walletSvc = service.NewWalletService()
}
