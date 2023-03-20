package repository

import (
	"github.com/xlalon/golee/internal/infra/repository/chain"
	"github.com/xlalon/golee/internal/infra/repository/deposit"
	"github.com/xlalon/golee/internal/infra/repository/wallet"
)

type Config struct {
	Chain   *chain.Config
	Deposit *deposit.Config
	Wallet  *wallet.Config
}

type Registry struct {
	chain   *chain.Dao
	deposit *deposit.Dao
	wallet  *wallet.Dao
}

func NewRegistry(conf *Config) *Registry {
	return &Registry{
		chain:   chain.NewDao(conf.Chain),
		deposit: deposit.NewDao(conf.Deposit),
		wallet:  wallet.NewDao(conf.Wallet),
	}
}

func (r *Registry) ChainRepository() *chain.Dao {
	return r.chain
}

func (r *Registry) DepositRepository() *deposit.Dao {
	return r.deposit
}

func (r *Registry) WalletRepository() *wallet.Dao {
	return r.wallet
}
