package deposit

import (
	"github.com/xlalon/golee/pkg/database/mysql"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type Deposit struct {
	mysql.Model

	Chain     string
	Asset     string
	TxHash    string
	Sender    string
	Receiver  string
	Memo      string
	Identity  string
	Amount    decimal.Decimal
	AmountRaw decimal.Decimal
	VOut      int64
	Status    string
	Comment   string

	Version int64
}

type IncomeCursor struct {
	mysql.Model

	ChainCode string
	Height    int64
	TxHash    string
	Address   string
	Label     string
	Index     int64
}
