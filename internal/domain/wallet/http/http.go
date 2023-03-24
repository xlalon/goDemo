package http

import (
	"github.com/xlalon/golee/internal/domain/wallet/conf"
	"github.com/xlalon/golee/internal/domain/wallet/service"
	"github.com/xlalon/golee/internal/onchain/x"
	"github.com/xlalon/golee/pkg/net/http/server"
)

var (
	_accountHandler = &accountHandler{}
)

var (
	walletService *service.WalletService
)

func Init(r *server.Engine, conf *conf.Config) error {
	initServices(conf)
	registerRouter(r)
	return nil
}

func registerRouter(r *server.Engine) {
	v1 := r.Group("/v1")

	rAccount := v1.Group("account")
	rAccount.POST("/new", _accountHandler.newAccount)
	rAccount.GET("/:address/detail", _accountHandler.getAccountDetail)
	rAccount.GET("/:address/balance", _accountHandler.getAccountBalance)
	rAccount.GET("/list", _accountHandler.getAccounts)
}

func initServices(conf *conf.Config) {
	_ = x.Init(conf.Chain)
	service.Init(conf)
	walletService = service.NewWalletService()
}
