package domain

type DepositRepository interface {
	Save(*Deposit) error
	GetDepositById(id int64) (*Deposit, error)
	GetDeposits() ([]*Deposit, error)
}
