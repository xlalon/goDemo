package onchain

type NodeInterface interface {
	GetNodeInfo() (*Node, error)
	GetLatestHeight() (int64, error)
}

type BlockInterface interface {
	GetBlockHash(height int64) (string, error)
	GetTxnByHash(txHash string) ([]*Transaction, error)
	ScanTxn(cursor *Cursor) ([]*Transaction, error)
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
