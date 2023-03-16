package model

import (
	"github.com/xlalon/golee/pkg/database/mysql"
)

type Asset struct {
	mysql.Model

	ChainCode string
	Code      string
	Name      string
	Identity  string
	Precision int64
	Status    string
}
