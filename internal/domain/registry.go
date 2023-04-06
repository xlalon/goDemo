package domain

import (
	"github.com/xlalon/golee/internal/domain/model/account"
	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/internal/domain/model/deposit"
	"github.com/xlalon/golee/internal/domain/service"
	"github.com/xlalon/golee/internal/onchain"
)

type Registry struct {
	AccountRepository account.AccountRepository
	ChainRepository   chainasset.ChainRepository
	DepositRepository deposit.DepositRepository

	OnChainSvc *onchain.Service
	AccountSvc *service.AccountService
	IncomeSvc  *service.Income
}

var (
	DomainRegistry *Registry
)

func Init(chainRepo chainasset.ChainRepository, depositRepo deposit.DepositRepository, accountRepo account.AccountRepository) {
	onChainSvc := onchain.NewService()

	DomainRegistry = &Registry{

		AccountRepository: accountRepo,
		ChainRepository:   chainRepo,
		DepositRepository: depositRepo,

		OnChainSvc: onChainSvc,
		AccountSvc: service.NewAccountService(accountRepo, chainRepo, onChainSvc),
		IncomeSvc:  service.NewIncome(accountRepo, chainRepo, depositRepo, onChainSvc),
	}
}
