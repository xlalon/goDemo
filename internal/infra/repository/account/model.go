package account

import (
	"github.com/xlalon/golee/pkg/database/mysql"
)

type Account struct {
	mysql.Model

	Chain   string
	Address string
	Label   string
	Memo    string
	Status  string
	Version int64
}
