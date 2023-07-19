package xchain

type WalletLabel string

const (
	WalletLabelDeposit WalletLabel = "DEPOSIT"
	WalletLabelHot     WalletLabel = "HOT"
)

type Wallet struct {
	Chain    Chain       `json:"chain"`
	Label    WalletLabel `json:"label"`
	Accounts []*Account  `json:"accounts"`
}

type Account struct {
	Address    Address
	privateKey []byte
	publicKey  []byte
	index      int64
}

type Address string
type Memo string

type AccountDTO struct {
	Address Address `json:"address"`
	Memo    Memo    `json:"memo"`
}
