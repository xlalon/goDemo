package chain

import (
	"context"

	"github.com/xlalon/golee/common/math/decimal"
)

type TransactionStatus int64

const (
	TransactionFailed  TransactionStatus = -1
	TransactionSuccess TransactionStatus = 0
	TransactionPending TransactionStatus = 1
	TransactionUnknown TransactionStatus = 2
)

type TransferTransaction struct {
	Chain         string
	TxHash        string
	VOut          int64
	Senders       []string
	Recipient     string
	Memo          string
	Crypto        Crypto
	Amount        decimal.Decimal
	Result        TransactionStatus
	Confirmations int64
	Height        int64
	Timestamp     int64
	Fee           *Fee
	Comment       interface{}
}

type Balance struct {
	Crypto Crypto
	Amount decimal.Decimal
}

type Fee struct {
	GasPrice decimal.Decimal
	GasLimit decimal.Decimal
	GasUsed  decimal.Decimal
	Crypto   Crypto
	Amount   decimal.Decimal
}

type Receipt struct {
	TxHash        string
	Result        TransactionStatus
	Confirmations int64
	Fee           *Fee
}

type BlockHeader struct {
	Height int64
	Hash   string
}

type Node interface {
	Version(context.Context) (string, error)
	Height(context.Context) (int64, error)
}

type Block interface {
	GetBlockHeader(ctx context.Context, heightOrHash interface{}) (*BlockHeader, error)
	GetBlockTransfers(ctx context.Context, heightOrHash interface{}) ([]*TransferTransaction, error)
	GetTxTransfers(ctx context.Context, txHash string) ([]*TransferTransaction, error)
	GetTxReceipt(ctx context.Context, txHash string) (*Receipt, error)
}

type WalletLabel string

const (
	Deposit WalletLabel = "DEPOSIT"
	Hot     WalletLabel = "HOT"
)

type Wallet interface {
	NewAccount(ctx context.Context) (Account, error)
	GetAccount(ctx context.Context, address string) (Account, error)
	Accounts(ctx context.Context) ([]Account, error)
	GetBalance(ctx context.Context, crypto *Crypto) (decimal.Decimal, error)
	Balances(ctx context.Context) ([]*Balance, error)
	GetTransfers(ctx context.Context, minHeight, maxHeight int64, in bool) ([]*TransferTransaction, error)
	Transfer(ctx context.Context, recipient, memo string, crypto Crypto, amount decimal.Decimal, data interface{}) (*Receipt, error)
}

type Account interface {
	Address(ctx context.Context) (string, error)
	Balance(ctx context.Context, crypto *Crypto) (decimal.Decimal, error)
	Balances(ctx context.Context) ([]*Balance, error)
	Nonce(ctx context.Context) (int64, error)
}

type Crypto interface {
	Id(ctx context.Context) string
	Symbol(ctx context.Context) string
	Name(ctx context.Context) string
	Precision(ctx context.Context) int64
}
