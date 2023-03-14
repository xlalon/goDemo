package http

import (
	"strconv"

	"github.com/xlalon/golee/pkg/net/http/server"
)

type depositHandler struct {
	server.Handler
}

func (dh *depositHandler) getDeposit(c *server.Context) {
	depId, _ := strconv.ParseInt(dh.Param(c, "id"), 10, 64)
	d, _ := depositSvc.GetDepositById(depId)
	dh.JSON(c, d)
}

func (dh *depositHandler) getDeposits(c *server.Context) {
	ds, _ := depositSvc.GetDeposits()
	dh.JSON(c, ds)
}
