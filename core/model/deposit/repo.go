package deposit

import "github.com/xlalon/golee/core/model/chain"

type Repo interface {
	Save(d *Deposit) error

	GetDepositById(id int64) (*Deposit, error)
	GetDepositsByTxHash(txHash string) ([]*Deposit, error)
	GetDepositsByChain(chain chain.Code) ([]*Deposit, error)
}
