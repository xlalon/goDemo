package chain

import "github.com/xlalon/golee/pkg/database/mysql"

type Chain struct {
	mysql.Model

	Code   string
	Name   string
	Status string
}

type Asset struct {
	mysql.Model

	ChainCode string
	Code      string
	Name      string
	Identity  string
	Precision int64
	Status    string
}
