package mysql

import (
	"gorm.io/gorm"
	"time"
)

type Config struct {
	DNS string `yaml:"dns"`
}

type Model struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
