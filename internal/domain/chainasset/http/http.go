package http

import (
	"github.com/xlalon/golee/internal/domain/chainasset/conf"
	"github.com/xlalon/golee/internal/domain/chainasset/service"
	"github.com/xlalon/golee/internal/onchain/x"
	"github.com/xlalon/golee/pkg/net/http/server"
)

var (
	_assetHandler = &assetHandler{}
	_chainHandler = &chainHandler{}
)

var (
	chainService *service.ChainService
	assetService *service.AssetService
)

func Init(r *server.Engine, conf *conf.Config) error {
	initServices(conf)
	registerRouter(r)
	return nil
}

func registerRouter(r *server.Engine) {
	v1 := r.Group("/v1")

	rChain := v1.Group("chain")
	rChain.GET("/:chain/height/latest", _chainHandler.getLatestHeight)
	rChain.GET("/list", _chainHandler.getChains)

	rAsset := v1.Group("asset")
	rAsset.GET("/list", _assetHandler.getAssets)

}

func initServices(conf *conf.Config) {
	_ = x.Init(conf.Chain)
	service.Init(conf)
	chainService = service.NewChainService()
	assetService = service.NewAssetService()

}
