package http

import (
	"strings"

	"github.com/xlalon/golee/pkg/net/http/server"
)

type chainHandler struct {
	server.Handler
}

func (ch *chainHandler) getLatestHeight(c *server.Context) {
	chainName := ch.Param(c, "chain")

	resp := chainSvc.GetLatestHeight(strings.ToUpper(chainName))

	ch.JSON(c, resp)
}
