package service

import (
	"github.com/xlalon/golee/internal/service/chain"
	"github.com/xlalon/golee/internal/service/deposit"
	"github.com/xlalon/golee/internal/service/deposit/domain"
	"github.com/xlalon/golee/internal/service/wallet"
)

type DepositService struct {
	depositSvc *deposit.Service
}

func NewDepositService(depositRepo domain.DepositRepository, chainSvc *chain.Service, walletSvc *wallet.Service) *DepositService {
	return &DepositService{
		depositSvc: deposit.NewService(depositRepo, chainSvc, walletSvc),
	}
}

func (d *DepositService) GetDepositById(id int64) (interface{}, error) {
	return d.depositSvc.GetDeposit(id)
}

func (d *DepositService) GetDeposits() (interface{}, error) {
	return d.depositSvc.GetDeposits()
}
