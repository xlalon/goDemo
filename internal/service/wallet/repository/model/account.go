package model

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID        int64          `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Version   int64          `json:"-"`

	Chain   string `json:"chain"`
	Address string `json:"address"`
	Label   string `json:"label"`
	Memo    string `json:"memo"`
	Status  string `json:"status"`
}
