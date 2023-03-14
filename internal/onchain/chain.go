package onchain

import "github.com/xlalon/golee/internal/onchain/conf"

type NodeInterface interface {
	GetNodeInfo() (*Node, error)
	GetLatestHeight() (int64, error)
}

type BlockInterface interface {
	GetBlockHash(height int64) (string, error)
	GetTxnByHash(txHash string) ([]*Transaction, error)
	ScanTxn(xxx interface{}) ([]*Transaction, error)
}

type WalletInterface interface {
	NewAccount(Label) (*Account, error)
	GetAccount(address string) (*Account, error)
	EstimateFee(*TransferDTO) (*Fee, error)
	Transfer(*TransferDTO) (*Receipt, error)
}

type Chainer interface {
	NodeInterface
	BlockInterface
	WalletInterface
}

type ScanTransfersByBlock interface {
	GetTxnByBlock(heightOrHash interface{}) ([]*Transaction, error)
}

type ScanTransfersByAccount interface {
	GetTxnByAccount(*Account) ([]*Transaction, error)
}

type Code string

type Chain struct {
	Code   Code              `json:"code"`
	Config *conf.ChainConfig `json:"config"`
}

func (*Chain) GetNodeInfo() (*Node, error) {
	return nil, nil
}

func (*Chain) GetBlockHash(height int64) (string, error) {
	_ = height
	return "", nil
}
