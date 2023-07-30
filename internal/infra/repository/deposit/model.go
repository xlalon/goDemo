package deposit

import (
	"github.com/xlalon/golee/pkg/database/mysql"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type Deposit struct {
	mysql.Model

	Chain    string
	TxHash   string
	VOut     int64
	Receiver string
	Memo     string
	Asset    string
	Amount   decimal.Decimal
	Sender   string
	Height   int64
	Comment  string
	Status   string
	Version  int64
}

type IncomeCursor struct {
	mysql.Model

	ChainCode   string
	Height      int64
	WalletLabel string
	Address     string
	TxHash      string
	Index       int64
}
