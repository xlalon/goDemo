package http

import (
	"github.com/xlalon/golee/pkg/net/http/server"
)

type depositHandler struct {
	server.Handler
}

func (dh *depositHandler) getDeposit(c *server.Context) {
	depId, _ := dh.QueryInt64(c, "id")
	d, _ := depositSvc.GetDepositById(depId)
	dh.JSON(c, d)
}

func (dh *depositHandler) getDeposits(c *server.Context) {
	ds, _ := depositSvc.GetDeposits()
	dh.JSON(c, ds)
}
