package model

import "gorm.io/gorm"

type Chain struct {
	gorm.Model

	Code string
	Name string

	Status string
}
