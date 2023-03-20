package wallet

import (
	"gorm.io/gorm"
	"time"

	"github.com/xlalon/golee/pkg/database/mysql"
)

type Account struct {
	mysql.Model

	Chain   string
	Address string
	Label   string
	Memo    string
	Status  string
	Version int64
}

type WalletHistory struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
