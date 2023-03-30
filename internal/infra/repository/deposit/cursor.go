package deposit

import (
	"time"

	"github.com/xlalon/golee/internal/domain/model/deposit"
	"github.com/xlalon/golee/pkg/database/mysql"
)

func (d *Dao) SaveIncomeCursor(cursor *deposit.IncomeCursor) error {
	var createdAt time.Time
	id := mysql.NextID()
	ic, err := d.getIncomeCursor(cursor.ChainCode())
	if err == nil && ic != nil {
		id = ic.ID
		createdAt = ic.CreatedAt
	}
	return d.db.Model(&IncomeCursor{Model: mysql.Model{ID: id}}).Save(&IncomeCursor{
		Model:     mysql.Model{ID: id, CreatedAt: createdAt},
		ChainCode: cursor.ChainCode(),
		Height:    cursor.Height(),
		TxHash:    cursor.TxHash(),
		Address:   cursor.Address(),
		Label:     cursor.Label(),
		Index:     cursor.Index(),
	}).Error
}

func (d *Dao) GetIncomeCursor(chainCode string) (*deposit.IncomeCursor, error) {
	ic, err := d.getIncomeCursor(chainCode)
	if err != nil {
		return nil, err
	}
	return deposit.NewIncomeCursor(&deposit.IncomeCursorDTO{
		ChainCode: chainCode,
		Height:    ic.Height,
		TxHash:    ic.TxHash,
		Address:   ic.Address,
		Label:     ic.Label,
		Index:     ic.Index,
	}), nil
}

func (d *Dao) getIncomeCursor(chainCode string) (*IncomeCursor, error) {
	ic := &IncomeCursor{}
	if err := d.db.First(ic, "chain_code = ?", chainCode).Error; err != nil {
		return nil, err
	}
	return ic, nil
}
