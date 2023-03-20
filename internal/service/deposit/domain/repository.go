package domain

import "github.com/xlalon/golee/pkg/database/mysql"

type DepositRepository interface {
	mysql.IdGeneratorRepository

	Save(*Deposit) error
	GetDepositById(id int64) (*Deposit, error)
	GetDeposits() ([]*Deposit, error)
}
