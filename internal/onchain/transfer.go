package onchain

import (
	"github.com/xlalon/golee/pkg/math/decimal"
)

type Status string

const (
	TransactionFailed  Status = "FAILED"
	TransactionPending        = "PENDING"
	TransactionSuccess        = "SUCCESS"
)

type Transaction struct {
	Chain     Code              `json:"chain"`
	TxHash    string            `json:"tx_hash"`
	VOut      int64             `json:"vout"`
	Status    TransactionStatus `json:"status"`
	Sender    string            `json:"sender"`
	Receiver  string            `json:"receiver"`
	Memo      string            `json:"memo"`
	Identity  string            `json:"identity"`
	Amount    decimal.Decimal   `json:"amount"`
	Height    int64             `json:"height"`
	Timestamp int64             `json:"timestamp"`
	Comment   interface{}       `json:"comment"`
}

type TransactionStatus struct {
	Result        Status `json:"result"`
	Confirmations int64  `json:"confirmations"`
}

type Fee struct {
	Identity string          `json:"identity"`
	Amount   decimal.Decimal `json:"amount"`
	Gas      int64           `json:"gas"`
	GasPrice decimal.Decimal `json:"gasPrice"`
}

type Receipt struct {
	TxHash string `json:"tx_hash"`
	Fee    Fee    `json:"fee"`
	Status Status `json:"status"`
	ErrLog string `json:"err_log"`
}

type TransferDTO struct {
	Sender, Receiver *Account
	Identity         string
	Amount           decimal.Decimal
	Fee              *Fee
	Extra            interface{}
}
