package onchain

import (
	"context"

	"github.com/xlalon/golee/internal/onchain/conf"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type NodeInterface interface {
	GetNodeInfo(context.Context) (*Node, error)
	GetLatestHeight(context.Context) (int64, error)
}

type BlockInterface interface {
	GetBlockHash(ctx context.Context, height int64) (string, error)
	GetTxnByHash(ctx context.Context, txHash string) ([]*Transaction, error)
	ScanTxn(context.Context, *Cursor) ([]*Transaction, error)
}

type WalletInterface interface {
	NewAccount(context.Context, Label) (*Account, error)
	GetAccount(ctx context.Context, address string) (*Account, error)
	GetBalance(ctx context.Context, account *Account, identity string) (decimal.Decimal, error)
	EstimateFee(context.Context, *TransferCommand) (*Fee, error)
	Transfer(context.Context, *TransferCommand) (*Receipt, error)
}

type Chainer interface {
	NodeInterface
	BlockInterface
	WalletInterface
}

type Chain struct {
	Code   Code
	Config *conf.ChainConfig
}

func (*Chain) GetNodeInfo(context.Context) (*Node, error) {
	return nil, nil
}

func (*Chain) GetBlockHash(context.Context, int64) (string, error) {
	return "", nil
}
