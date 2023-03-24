package service

import (
	"github.com/xlalon/golee/internal/domain/deposit/model"
)

type DepositService struct {
	domainRegistry *Registry
}

func NewDepositService() *DepositService {
	return &DepositService{DomainRegistry}
}

func (d *DepositService) GetDepositById(id int64) (interface{}, error) {
	return d.depositDMToDTO(d.domainRegistry.depositRepository.GetDepositById(id))
}

func (d *DepositService) GetDeposits(page, limit int64) (interface{}, error) {
	return d.depositsDMToDTOs(d.domainRegistry.depositRepository.GetDeposits(page, limit))
}

func (d *DepositService) depositDMToDTO(dep *model.Deposit, err error) (*model.DepositDTO, error) {
	if err != nil {
		return nil, err
	}
	return dep.ToDepositDTO(), nil
}

func (d *DepositService) depositsDMToDTOs(deps []*model.Deposit, err error) ([]*model.DepositDTO, error) {
	if err != nil {
		return nil, err
	}
	var depsDTOs []*model.DepositDTO
	for _, dep := range deps {
		depsDTOs = append(depsDTOs, dep.ToDepositDTO())
	}
	return depsDTOs, nil
}
