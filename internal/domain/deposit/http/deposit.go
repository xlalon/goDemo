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
	d, _ := depositService.GetDepositById(depId)
	dh.JSON(c, d)
}

func (dh *depositHandler) getDeposits(c *server.Context) {
	page, ok := dh.QueryInt64(c, "page")
	if !ok {
		page = 1
	}
	limit, ok := dh.QueryInt64(c, "limit")
	if !ok {
		limit = 50
	}
	ds, _ := depositService.GetDeposits(page, limit)
	dh.JSON(c, ds)
}
