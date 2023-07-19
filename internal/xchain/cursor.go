package xchain

type Direction string

type Cursor struct {
	Chain  Chain
	Height int64
	// scan by account
	WalletLabel WalletLabel
	Address     Address
	TxHash      string
	Index       int64
}

func NewCursor(chain Chain, height int64, walletLabel string, address, txHash string, index int64) *Cursor {
	_walletLabel := WalletLabel(walletLabel)
	if _walletLabel != WalletLabelDeposit && _walletLabel != WalletLabelHot {
		return nil
	}
	return &Cursor{
		Chain:       chain,
		Height:      height,
		WalletLabel: _walletLabel,
		Address:     Address(address),
		TxHash:      txHash,
		Index:       index,
	}
}
