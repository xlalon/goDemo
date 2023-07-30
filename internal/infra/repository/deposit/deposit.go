package deposit

import (
	"time"

	"github.com/xlalon/golee/internal/domain/model/deposit"
	"github.com/xlalon/golee/pkg/database/mysql"
)

func (d *Dao) Save(dep *deposit.Deposit) error {
	var createdAt time.Time
	var version int64
	depositDB, err := d.getDepositById(dep.Id())
	if err == nil && depositDB != nil {
		createdAt = depositDB.CreatedAt
		version = depositDB.Version + 1
	}
	d.db.Model(&Deposit{Model: mysql.Model{ID: dep.Id()}}).Save(&Deposit{
		Model: mysql.Model{
			ID:        dep.Id(),
			CreatedAt: createdAt,
		},
		Chain:    dep.Chain(),
		TxHash:   dep.TxHash(),
		VOut:     dep.VOut(),
		Receiver: dep.Receiver(),
		Memo:     dep.Memo(),
		Asset:    string(dep.Asset()),
		Amount:   dep.Amount(),
		Sender:   dep.Sender(),
		Height:   dep.Height(),
		Comment:  dep.Comment(),
		Status:   string(dep.Status()),
		Version:  version,
	})
	return nil
}

func (d *Dao) GetDepositById(id int64) (*deposit.Deposit, error) {
	return d.depositDbToDomain(d.getDepositById(id))
}

func (d *Dao) GetDeposits(page, limit int64) ([]*deposit.Deposit, error) {
	return d.depositsDbToDomain(d.getDeposits(page, limit))
}

func (d *Dao) getDepositById(id int64) (*Deposit, error) {
	depDB := Deposit{}
	if err := d.db.First(&depDB, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &depDB, nil
}

func (d *Dao) getDeposits(page, limit int64) ([]Deposit, error) {
	var depsDB []Deposit
	offset := (page - 1) * limit
	if offset <= 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 50
	}
	if err := d.db.Offset(int(offset)).Limit(int(limit)).Find(&depsDB).Error; err != nil {
		return nil, err
	}
	return depsDB, nil
}

func (d *Dao) depositsDbToDomain(depsDB []Deposit, err error) ([]*deposit.Deposit, error) {
	if err != nil || depsDB == nil || len(depsDB) == 0 {
		return nil, err
	}
	var deps []*deposit.Deposit
	for _, depDB := range depsDB {
		if depDM, _ := d.depositDbToDomain(&depDB, nil); depDM != nil {
			deps = append(deps, depDM)
		}
	}
	return deps, nil
}

func (d *Dao) depositDbToDomain(dep *Deposit, err error) (*deposit.Deposit, error) {
	if err != nil || dep == nil {
		return nil, err
	}
	return deposit.DepositFactory(&deposit.DepositDTO{
		Id:       dep.ID,
		Chain:    dep.Chain,
		TxHash:   dep.TxHash,
		VOut:     dep.VOut,
		Receiver: dep.Receiver,
		Memo:     dep.Memo,
		Asset:    dep.Asset,
		Amount:   dep.Amount,
		Sender:   dep.Sender,
		Height:   dep.Height,
		Comment:  dep.Comment,
		Status:   deposit.Status(dep.Status),
	}), nil
}
