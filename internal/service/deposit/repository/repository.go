package repository

import (
	"github.com/xlalon/golee/internal/service/deposit/domain"
)

type DepositRepository interface {
	CreateDeposit(dto *domain.DepositDTO) error
	GetDepositById(id int64) (*domain.Deposit, error)
	GetDeposits() ([]*domain.Deposit, error)
}
