package http

import (
	"github.com/xlalon/golee/internal/app/conf"
	"github.com/xlalon/golee/internal/app/service"
	"github.com/xlalon/golee/internal/onchain/x"
	"github.com/xlalon/golee/pkg/net/http/server"
)

var (
	_assetHandler   = &assetHandler{}
	_chainHandler   = &chainHandler{}
	_depositHandler = &depositHandler{}
	_walletHandler  = &walletHandler{}
)

var (
	assetSvc   *service.AssetService
	chainSvc   *service.ChainService
	depositSvc *service.DepositService
	walletSvc  *service.WalletService
)

func Init(r *server.Engine, conf *conf.Config) error {
	initServices(conf)
	registerRouter(r)
	return nil
}

func registerRouter(r *server.Engine) {
	v1 := r.Group("/v1")

	rAsset := v1.Group("asset")
	rAsset.GET("/list", _assetHandler.getAssets)

	rChain := v1.Group("chain")
	rChain.GET("/:chain/height/latest", _chainHandler.getLatestHeight)
	rChain.GET("/list", _chainHandler.getChains)

	rDeposit := v1.Group("deposit")
	rDeposit.GET("/:id", _depositHandler.getDeposit)
	rDeposit.GET("/list", _depositHandler.getDeposits)

	rAccount := v1.Group("account")
	rAccount.POST("/new", _walletHandler.newAccount)
	rAccount.GET("/:address/detail", _walletHandler.getAccountDetail)
	rAccount.GET("/:address/balance", _walletHandler.getAccountBalance)
	rAccount.GET("/list", _walletHandler.getAccounts)

}

func initServices(conf *conf.Config) {
	x.Init(conf.Chain)
	service.Init(conf)
	chainSvc = service.NewChainService()
	assetSvc = service.NewAssetService()
	depositSvc = service.NewDepositService()
	walletSvc = service.NewWalletService()
}
