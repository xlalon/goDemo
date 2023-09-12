package chain

import (
	"context"

	"github.com/xlalon/golee/x/conf"
)

type Chainer interface {
	Node
	Block

	GetConfig(ctx context.Context) (*conf.Config, error)
	GetWallet(ctx context.Context, walletLabel WalletLabel) (Wallet, error)
	GetAccount(ctx context.Context, address string) (Account, error)
}

type Chain struct {
	Code string

	Config *conf.Config
}

func (c *Chain) Version(context.Context) (string, error) {
	return "", nil
}

func (c *Chain) Height(context.Context) (int64, error) {
	return 0, nil
}

func (c *Chain) GetBlockHeader(ctx context.Context, heightOrHash interface{}) (*BlockHeader, error) {
	return nil, nil
}

func (c *Chain) GetBlockTransfers(ctx context.Context, heightOrHash interface{}) ([]*TransferTransaction, error) {
	return nil, nil
}

func (c *Chain) GetTxTransfers(ctx context.Context, txHash string) ([]*TransferTransaction, error) {
	return nil, nil
}

func (c *Chain) GetTxReceipt(ctx context.Context, txHash string) (*Receipt, error) {
	return nil, nil
}

func (c *Chain) GetConfig(ctx context.Context) (*conf.Config, error) {
	return nil, nil
}

func (c *Chain) GetWallet(ctx context.Context, walletLabel WalletLabel) (Wallet, error) {
	return nil, nil
}

func (c *Chain) GetAccount(ctx context.Context, address string) (Account, error) {
	return nil, nil
}
