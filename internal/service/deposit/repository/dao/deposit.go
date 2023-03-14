package dao

import (
	"github.com/xlalon/golee/internal/service/deposit/domain"
	"github.com/xlalon/golee/internal/service/deposit/repository/model"
)

func (d *Dao) CreateDeposit(dto *domain.DepositDTO) error {
	return d.db.Create(&model.Deposit{
		Chain:     dto.Chain,
		Asset:     dto.Asset,
		TxHash:    dto.TxHash,
		Sender:    dto.Sender,
		Receiver:  dto.Receiver,
		Memo:      dto.Memo,
		Identity:  dto.Identity,
		Amount:    dto.Amount,
		AmountRaw: dto.AmountRaw,
		VOut:      dto.VOut,
		Status:    string(domain.DepositStatusPending),
		Comment:   "",
	}).Error
}

func (d *Dao) GetDeposits() ([]*domain.Deposit, error) {
	var depsDB []model.Deposit
	if err := d.db.Find(&depsDB).Error; err != nil {
		return nil, err
	}
	var deps []*domain.Deposit
	for _, dep := range depsDB {
		deps = append(deps, d.depositDbToDomain(&dep))
	}
	return deps, nil
}

func (d *Dao) GetDepositById(id int64) (*domain.Deposit, error) {
	depDB := &model.Deposit{}
	if err := d.db.First(depDB, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return d.depositDbToDomain(depDB), nil
}

func (d *Dao) depositDbToDomain(dep *model.Deposit) *domain.Deposit {
	return &domain.Deposit{
		Id: int64(dep.ID),
		DepositItem: *domain.NewDepositItem(&domain.DepositDTO{
			Chain:     dep.Chain,
			Asset:     dep.Asset,
			TxHash:    dep.TxHash,
			Sender:    dep.Sender,
			Receiver:  dep.Receiver,
			Memo:      dep.Memo,
			Identity:  dep.Identity,
			Amount:    dep.Amount,
			AmountRaw: dep.AmountRaw,
			VOut:      dep.VOut,
		}),
		Status: domain.Status(dep.Status),
	}
}
