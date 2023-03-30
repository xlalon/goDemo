package deposit

import "github.com/xlalon/golee/pkg/database/mysql"

type DepositRepository interface {
	mysql.IdGeneratorRepository

	Save(*Deposit) error
	GetDepositById(id int64) (*Deposit, error)
	GetDeposits(page, limit int64) ([]*Deposit, error)

	SaveIncomeCursor(cursor *IncomeCursor) error
	GetIncomeCursor(chainCode string) (*IncomeCursor, error)
}
