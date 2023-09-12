package mysql

import (
	"gorm.io/gorm"
	"time"
)

type Config struct {
	DSN string `yaml:"dsn"`
}

type Model struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type IdGeneratorRepository interface {
	NextId() int64
}
