package deposit

type IncomeCursor struct {
	chainCode string
	height    int64
	WalletIncomeCursor
	AccountIncomeCursor
}

type WalletIncomeCursor struct {
	label  string
	txHash string
	index  int64
}

type AccountIncomeCursor struct {
	address string
	txHash  string
	index   int64
}
