package http

import (
	"github.com/xlalon/golee/internal/admin/conf"
	"github.com/xlalon/golee/internal/admin/service"
	"github.com/xlalon/golee/internal/onchain/x"
	"github.com/xlalon/golee/pkg/net/http/server"
)

var (
	chainSvc   *service.ChainService
	depositSvc *service.DepositService
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
	chainSvc = service.NewChainService(conf)
	depositSvc = service.NewDepositService(conf)
}
