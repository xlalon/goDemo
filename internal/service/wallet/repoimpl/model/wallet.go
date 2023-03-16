package model

import (
	"gorm.io/gorm"
	"time"
)

type WalletHistory struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
