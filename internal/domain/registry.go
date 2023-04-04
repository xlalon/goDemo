package domain

import (
	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/internal/domain/model/deposit"
	"github.com/xlalon/golee/internal/domain/model/wallet"
	"github.com/xlalon/golee/internal/onchain"
)

type Registry struct {
	ChainRepository   chainasset.ChainRepository
	DepositRepository deposit.DepositRepository
	WalletRepository  wallet.WalletRepository

	OnChainSvc *onchain.Service
}

var (
	DomainRegistry *Registry
)

func Init(chainRepo chainasset.ChainRepository, depositRepo deposit.DepositRepository, walletRepo wallet.WalletRepository) {

	DomainRegistry = &Registry{

		ChainRepository:   chainRepo,
		DepositRepository: depositRepo,
		WalletRepository:  walletRepo,

		OnChainSvc: onchain.NewService(),
	}
}
