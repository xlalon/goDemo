package deposit

import (
	"github.com/xlalon/golee/internal/service/deposit/domain"
	"time"

	"github.com/xlalon/golee/pkg/database/mysql"
)

func (d *Dao) Save(dep *domain.Deposit) error {
	var createAt = time.Time{}
	var version int64
	depositDB, err := d.getDepositById(dep.GetId())
	if err == nil && depositDB != nil {
		createAt = depositDB.CreatedAt
		version = depositDB.Version + 1
	}
	d.db.Model(&Deposit{Model: mysql.Model{ID: dep.GetId()}}).Save(&Deposit{
		Model: mysql.Model{
			ID:        dep.GetId(),
			CreatedAt: createAt,
		},
		Chain:     dep.GetChain(),
		Asset:     dep.GetAsset(),
		TxHash:    dep.GetTxHash(),
		Sender:    dep.GetSender(),
		Receiver:  dep.GetReceiver(),
		Memo:      dep.GetMemo(),
		Identity:  dep.GetIdentity(),
		Amount:    dep.GetAmount(),
		AmountRaw: dep.GetAmountRaw(),
		VOut:      dep.GetVOut(),
		Status:    string(dep.GetStatus()),
		Comment:   "",
		Version:   version,
	})
	return nil
}

func (d *Dao) GetDepositById(id int64) (*domain.Deposit, error) {
	depDB, err := d.getDepositById(id)
	if err != nil {
		return nil, err
	}
	return d.depositDbToDomain(depDB), nil
}

func (d *Dao) GetDeposits() ([]*domain.Deposit, error) {
	var depsDB []Deposit
	if err := d.db.Find(&depsDB).Error; err != nil {
		return nil, err
	}
	var deps []*domain.Deposit
	for _, dep := range depsDB {
		deps = append(deps, d.depositDbToDomain(&dep))
	}
	return deps, nil
}

func (d *Dao) getDepositById(id int64) (*Deposit, error) {
	depDB := &Deposit{}
	if err := d.db.First(depDB, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return depDB, nil
}

func (d *Dao) depositDbToDomain(dep *Deposit) *domain.Deposit {
	return domain.DepositFactory(&domain.DepositDTO{
		Id:        dep.ID,
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
		Status:    domain.Status(dep.Status),
	})

}
