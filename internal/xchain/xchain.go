package xchain

import (
	"context"

	"github.com/xlalon/golee/internal/xchain/conf"
)

type NodeInterface interface {
	GetNodeVersion(context.Context) (string, error)
	GetLatestHeight(context.Context) (int64, error)
}

type BlockInterface interface {
	GetBlockHash(ctx context.Context, height int64) (string, error)
	GetTransfersByHash(ctx context.Context, txHash string) ([]*Transfer, error)
	ScanTransfers(context.Context, *Cursor) ([]*Transfer, error)
}

type WalletInterface interface {
	NewAccount(context.Context, WalletLabel) (*AccountDTO, error)
	GetWalletBalance(context.Context, WalletLabel, Identity) (Coin, error)
	GetAccountBalance(context.Context, Address, Identity) (Coin, error)
	EstimateFee(context.Context, *TransferCommand) (*Fee, error)
	Transfer(context.Context, *TransferCommand) (*Receipt, error)
}

type Chainer interface {
	NodeInterface
	BlockInterface
	WalletInterface
}

type X struct {
	Code   Chain
	Config *conf.ChainConfig
}

func (*X) GetNodeVersion(context.Context) (string, error) {
	return "", nil
}

func (*X) GetBlockHash(context.Context, int64) (string, error) {
	return "", nil
}
