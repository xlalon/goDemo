package model

import (
	"gorm.io/gorm"
	"time"
)

type WalletHistory struct {
	ID        int64          `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
