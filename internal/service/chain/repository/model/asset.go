package model

import "gorm.io/gorm"

type Asset struct {
	gorm.Model

	ChainCode string

	Code      string
	Name      string
	Identity  string
	Precision int64

	Status string
}
