package wallet

import (
	"github.com/xlalon/golee/pkg/database/mysql"
)

type AccountRepository interface {
	mysql.IdGeneratorRepository

	Save(acct *Account) error
	GetAccountByChainAddress(chain, address string) (*Account, error)
	GetAccountsByChain(chain string) ([]*Account, error)
	GetAccountsByChainAddresses(chain string, addresses []string) ([]*Account, error)
}
