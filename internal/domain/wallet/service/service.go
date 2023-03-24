package service

import (
	"github.com/xlalon/golee/internal/domain/wallet/conf"
	"github.com/xlalon/golee/internal/domain/wallet/model"
	"github.com/xlalon/golee/internal/infra/repository/wallet"
	"github.com/xlalon/golee/internal/onchain"
)

type Registry struct {
	walletRepository model.WalletRepository
	onChainService   *onchain.Service
}

var (
	DomainRegistry *Registry
)

func Init(conf *conf.Config) {
	walletRepo := wallet.NewDao(&wallet.Config{
		Mysql: conf.Mysql,
		Redis: conf.Redis,
	})
	DomainRegistry = &Registry{
		walletRepository: walletRepo,
		onChainService:   onchain.NewService(),
	}
}
