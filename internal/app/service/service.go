package service

import (
	"github.com/xlalon/golee/internal/app/conf"
	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/internal/domain/model/deposit"
	"github.com/xlalon/golee/internal/domain/model/wallet"
	chainRepo "github.com/xlalon/golee/internal/infra/repository/chainasset"
	depositRepo "github.com/xlalon/golee/internal/infra/repository/deposit"
	walletRepo "github.com/xlalon/golee/internal/infra/repository/wallet"
	"github.com/xlalon/golee/internal/onchain"
)

type Registry struct {
	chainRepository   chainasset.ChainRepository
	depositRepository deposit.DepositRepository
	walletRepository  wallet.WalletRepository

	chainAssetSvc *chainasset.Service

	onChainService *onchain.Service
}

var (
	DomainRegistry *Registry
)

func Init(conf *conf.Config) {
	chainRepository := chainRepo.NewDao(&chainRepo.Config{
		Mysql: conf.Mysql,
		Redis: conf.Redis,
	})
	depositRepository := depositRepo.NewDao(&depositRepo.Config{
		Mysql: conf.Mysql,
		Redis: conf.Redis,
	})
	walletRepository := walletRepo.NewDao(&walletRepo.Config{
		Mysql: conf.Mysql,
		Redis: conf.Redis,
	})

	DomainRegistry = &Registry{

		chainRepository:   chainRepository,
		depositRepository: depositRepository,
		walletRepository:  walletRepository,

		chainAssetSvc: chainasset.NewService(chainRepository),

		onChainService: onchain.NewService(),
	}
}
