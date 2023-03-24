package service

import (
	"github.com/xlalon/golee/internal/domain/deposit/conf"
	"github.com/xlalon/golee/internal/domain/deposit/model"
	"github.com/xlalon/golee/internal/infra/repository/deposit"
	"github.com/xlalon/golee/internal/onchain"
)

type Registry struct {
	depositRepository model.DepositRepository
	onChainService    *onchain.Service
}

var (
	DomainRegistry *Registry
)

func Init(conf *conf.Config) {
	depositRepo := deposit.NewDao(&deposit.Config{
		Mysql: conf.Mysql,
		Redis: conf.Redis,
	})
	DomainRegistry = &Registry{
		depositRepository: depositRepo,
		onChainService:    onchain.NewService(),
	}
}
