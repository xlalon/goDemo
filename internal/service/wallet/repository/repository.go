package repository

import (
	"github.com/xlalon/golee/internal/service/wallet/domain"
)

type WalletRepository interface {
	NewAccount(acct *domain.Account) error
	GetAccountByChainAddress(chain, address string) (*domain.Account, error)
	GetAccountsByChain(chain string) ([]*domain.Account, error)
	GetAccountsByChainAddresses(chain string, addresses []string) ([]*domain.Account, error)
}
