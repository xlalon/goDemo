package model

import "github.com/xlalon/golee/pkg/database/mysql"

type Chain struct {
	mysql.Model

	Code   string
	Name   string
	Status string
}
