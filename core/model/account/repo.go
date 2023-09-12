package account

type Repo interface {
	Save(acct *Account) error

	GetAccountById(id int64) (*Account, error)
	GetAccountByAddressAndMemo(address, memo string) (*Account, error)
	GetAccountsByAddressAndMemos(addressAndMemos [][2]string) ([]*Account, error)
	GetAccounts() ([]*Account, error)
}
