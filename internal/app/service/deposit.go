package service

import (
	"github.com/xlalon/golee/internal/domain"
	"github.com/xlalon/golee/internal/domain/model/deposit"
)

type DepositService struct {
	Service
}

func NewDepositService() *DepositService {
	return &DepositService{Service{DomainRegistry: domain.DomainRegistry}}
}

func (d *DepositService) GetDepositById(id int64) (interface{}, error) {
	return d.depositDMToDTO(d.DomainRegistry.DepositRepository.GetDepositById(id))
}

func (d *DepositService) GetDeposits(page, limit int64) (interface{}, error) {
	return d.depositsDMToDTOs(d.DomainRegistry.DepositRepository.GetDeposits(page, limit))
}

func (d *DepositService) depositDMToDTO(dep *deposit.Deposit, err error) (*deposit.DepositDTO, error) {
	if err != nil {
		return nil, err
	}
	return dep.ToDepositDTO(), nil
}

func (d *DepositService) depositsDMToDTOs(deps []*deposit.Deposit, err error) ([]*deposit.DepositDTO, error) {
	if err != nil {
		return nil, err
	}
	var depsDTOs []*deposit.DepositDTO
	for _, dep := range deps {
		depsDTOs = append(depsDTOs, dep.ToDepositDTO())
	}
	return depsDTOs, nil
}
