package deposit

import (
	"time"

	"github.com/xlalon/golee/internal/domain/model/deposit"
	"github.com/xlalon/golee/pkg/database/mysql"
)

func (d *Dao) SaveIncomeCursor(cursor *deposit.IncomeCursor) error {
	var createdAt time.Time
	id := mysql.NextID()
	ic, err := d.getIncomeCursor(cursor.ChainCode(), cursor.Address(), cursor.Label())
	if err == nil && ic != nil {
		id = ic.ID
		createdAt = ic.CreatedAt
	}
	return d.db.Model(&IncomeCursor{Model: mysql.Model{ID: id}}).Save(&IncomeCursor{
		Model: mysql.Model{
			ID:        id,
			CreatedAt: createdAt,
		},
		ChainCode: cursor.ChainCode(),
		Height:    cursor.Height(),
		Address:   cursor.Address(),
		Label:     cursor.Label(),
		TxHash:    cursor.TxHash(),
		Direction: cursor.Direction(),
		Index:     cursor.Index(),
	}).Error
}

func (d *Dao) GetIncomeCursor(chainCode, address, label string) (*deposit.IncomeCursor, error) {
	ic, err := d.getIncomeCursor(chainCode, address, label)
	if err != nil {
		return nil, err
	}
	return deposit.NewIncomeCursor(
		chainCode,
		ic.Height,
		ic.Address,
		ic.Label,
		ic.TxHash,
		ic.Direction,
		ic.Index,
	), nil
}

func (d *Dao) getIncomeCursor(chainCode, address, label string) (*IncomeCursor, error) {
	ic := &IncomeCursor{}
	if err := d.db.First(ic, "chain_code = ? AND address = ? AND label = ?", chainCode, address, label).Error; err != nil {
		return nil, err
	}
	return ic, nil
}
