package onchain

import (
	"github.com/xlalon/golee/pkg/math/decimal"
)

type Status string

const (
	TxnFailed  Status = "FAILED"
	TxnPending        = "PENDING"
	TxnSuccess        = "SUCCESS"
)

type TxnId struct {
	Chain  Code
	TxHash string
	VOut   int64
}

type Receiver struct {
	Address string
	Memo    string
}

type CoinValue struct {
	Identity string
	Amount   decimal.Decimal
}

type TxnStatus struct {
	Result        Status
	Confirmations int64
}

type Transaction struct {
	TxnId     TxnId
	Receiver  Receiver
	CoinValue CoinValue
	Status    TxnStatus
	Sender    string
	Height    int64
	Timestamp int64
	Comment   string
}

type Fee struct {
	CoinValue CoinValue
	GasPrice  decimal.Decimal
	GasLimit  int64
}

type Receipt struct {
	TxHash string
	Fee    *Fee
	Status Status
	ErrLog string
}

type TransferCommand struct {
	Sender    *Account
	Receiver  Receiver
	CoinValue CoinValue
	Fee       *Fee
	Extra     map[string]interface{}
}
