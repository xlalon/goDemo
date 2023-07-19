package domain

import (
	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/internal/domain/model/deposit"
	"github.com/xlalon/golee/internal/domain/model/wallet"
	"github.com/xlalon/golee/internal/domain/service"
	"github.com/xlalon/golee/internal/xchain"
)

type Registry struct {
	AccountRepository wallet.AccountRepository
	ChainRepository   chainasset.ChainRepository
	DepositRepository deposit.DepositRepository

	OnChainSvc *xchain.Service
	AccountSvc *service.AccountService
	IncomeSvc  *service.Income
}

var (
	DomainRegistry *Registry
)

func Init(chainRepo chainasset.ChainRepository, depositRepo deposit.DepositRepository, accountRepo wallet.AccountRepository) {
	onChainSvc := xchain.NewService()

	DomainRegistry = &Registry{

		AccountRepository: accountRepo,
		ChainRepository:   chainRepo,
		DepositRepository: depositRepo,

		OnChainSvc: onChainSvc,
		AccountSvc: service.NewAccountService(accountRepo, chainRepo, onChainSvc),
		IncomeSvc:  service.NewIncome(accountRepo, chainRepo, depositRepo, onChainSvc),
	}
}
