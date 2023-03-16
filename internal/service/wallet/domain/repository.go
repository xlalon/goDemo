package domain

type WalletRepository interface {
	Save(acct *Account) error
	GetAccountByChainAddress(chain, address string) (*Account, error)
	GetAccountsByChain(chain string) ([]*Account, error)
	GetAccountsByChainAddresses(chain string, addresses []string) ([]*Account, error)
}
