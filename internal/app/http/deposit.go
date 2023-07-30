package http

import (
	"strconv"

	"github.com/xlalon/golee/pkg/net/http/server"
)

type depositHandler struct {
	server.Handler
}

func (dh *depositHandler) getDeposit(ctx *server.Context) {
	depId, _ := strconv.ParseInt(dh.Param(ctx, "id"), 10, 64)
	d, _ := depositSvc.GetDepositById(depId)
	dh.JSON(ctx, d)
}

func (dh *depositHandler) getDeposits(ctx *server.Context) {
	page, ok := dh.QueryInt64(ctx, "page")
	if !ok {
		page = 1
	}
	limit, ok := dh.QueryInt64(ctx, "limit")
	if !ok {
		limit = 50
	}
	ds, _ := depositSvc.GetDeposits(page, limit)
	dh.JSON(ctx, ds)
}
