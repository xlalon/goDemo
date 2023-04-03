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
	Gas       int64
	GasPrice  decimal.Decimal
}

type Receipt struct {
	TxHash string
	Fee    Fee
	Status Status
	ErrLog string
}

type TransferDTO struct {
	Sender, Receiver *Account
	Identity         string
	Amount           decimal.Decimal
	Fee              *Fee
	Extra            interface{}
}
