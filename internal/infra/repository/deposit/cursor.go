package deposit

import (
	"github.com/xlalon/golee/internal/domain/model/deposit"
)

func (d *Dao) SaveIncomeCursor(cursor *deposit.IncomeCursor) error {
	return nil
}

func (d *Dao) GetIncomeCursor(chainCode, address, label string) (*deposit.IncomeCursor, error) {
	return &deposit.IncomeCursor{}, nil
}

func (d *Dao) getIncomeCursor(chainCode, address, label string) (*IncomeCursor, error) {
	return &IncomeCursor{}, nil
}
