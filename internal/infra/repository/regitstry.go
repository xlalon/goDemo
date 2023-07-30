package repository

import (
	"github.com/xlalon/golee/internal/infra/repository/account"
	"github.com/xlalon/golee/internal/infra/repository/chainasset"
	"github.com/xlalon/golee/internal/infra/repository/deposit"
)

type Config struct {
	Chain   *chainasset.Config
	Deposit *deposit.Config
	Account *account.Config
}

type Registry struct {
	chain   *chainasset.Dao
	deposit *deposit.Dao
	account *account.Dao
}

func NewRegistry(conf *Config) *Registry {
	return &Registry{
		chain:   chainasset.NewDao(conf.Chain),
		deposit: deposit.NewDao(conf.Deposit),
		account: account.NewDao(conf.Account),
	}
}

func (r *Registry) ChainRepository() *chainasset.Dao {
	return r.chain
}

func (r *Registry) DepositRepository() *deposit.Dao {
	return r.deposit
}

func (r *Registry) AccountRepository() *account.Dao {
	return r.account
}
