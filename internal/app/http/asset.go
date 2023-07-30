package http

import (
	"github.com/xlalon/golee/pkg/net/http/server"
)

type assetHandler struct {
	server.Handler
}

func (ah *assetHandler) getAssets(ctx *server.Context) {

	resp, _ := assetSvc.GetAssets()

	ah.JSON(ctx, resp)
}
