package http

import (
	"github.com/xlalon/golee/internal/domain/deposit/conf"
	"github.com/xlalon/golee/internal/domain/deposit/service"
	"github.com/xlalon/golee/internal/onchain/x"
	"github.com/xlalon/golee/pkg/net/http/server"
)

var (
	_depositHandler = &depositHandler{}
)

var (
	depositService *service.DepositService
)

func Init(r *server.Engine, conf *conf.Config) error {
	initServices(conf)
	registerRouter(r)
	return nil
}

func registerRouter(r *server.Engine) {
	v1 := r.Group("/v1")

	rDeposit := v1.Group("deposit")
	rDeposit.GET("/:id", _depositHandler.getDeposit)
	rDeposit.GET("/list", _depositHandler.getDeposits)

}

func initServices(conf *conf.Config) {
	_ = x.Init(conf.Chain)
	service.Init(conf)
	depositService = service.NewDepositService()
}
