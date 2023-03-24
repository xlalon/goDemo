package service

import (
	"github.com/xlalon/golee/internal/domain/chainasset/conf"
	"github.com/xlalon/golee/internal/domain/chainasset/model"
	"github.com/xlalon/golee/internal/infra/repository/chainasset"
	"github.com/xlalon/golee/internal/onchain"
)

type Registry struct {
	chainRepository   model.ChainRepository
	chainAssetService *model.Service
	onChainService    *onchain.Service
}

var (
	DomainRegistry *Registry
)

func Init(conf *conf.Config) {
	chainRepo := chainasset.NewDao(&chainasset.Config{
		Mysql: conf.Mysql,
		Redis: conf.Redis,
	})
	DomainRegistry = &Registry{
		chainRepository:   chainRepo,
		chainAssetService: model.NewService(chainRepo),
		onChainService:    onchain.NewService(),
	}
}
