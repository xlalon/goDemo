package onchain

import "context"

type NodeInterface interface {
	GetNodeInfo(ctx context.Context) (*Node, error)
	GetLatestHeight(ctx context.Context) (int64, error)
}

type BlockInterface interface {
	GetBlockHash(ctx context.Context, height int64) (string, error)
	GetTxnByHash(ctx context.Context, txHash string) ([]*Transaction, error)
	ScanTxn(ctx context.Context, cursor *Cursor) ([]*Transaction, error)
}

type WalletInterface interface {
	NewAccount(context.Context, Label) (*Account, error)
	GetAccount(ctx context.Context, address string) (*Account, error)
	EstimateFee(context.Context, *TransferCommand) (*Fee, error)
	Transfer(context.Context, *TransferCommand) (*Receipt, error)
}

type Chainer interface {
	NodeInterface
	BlockInterface
	WalletInterface
}
