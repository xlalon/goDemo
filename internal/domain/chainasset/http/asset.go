package http

import (
	"github.com/xlalon/golee/pkg/net/http/server"
)

type assetHandler struct {
	server.Handler
}

func (ah *assetHandler) getAssets(c *server.Context) {

	resp, _ := assetService.GetAssets()

	ah.JSON(c, resp)
}
