package http

import (
	"strings"

	"github.com/xlalon/golee/pkg/net/http/server"
)

type chainHandler struct {
	server.Handler
}

func (ch *chainHandler) getLatestHeight(ctx *server.Context) {
	chainName := ch.Param(ctx, "chain")

	resp, _ := chainSvc.GetChainLatestHeight(ctx, strings.ToUpper(chainName))

	ch.JSON(ctx, resp)
}

func (ch *chainHandler) getChains(ctx *server.Context) {

	resp, _ := chainSvc.GetChains()

	ch.JSON(ctx, resp)
}
