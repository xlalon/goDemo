package domain

type Account struct {
	Chain    string     `json:"chain"`
	Address  string     `json:"address"`
	Label    string     `json:"label"`
	Memo     string     `json:"memo"`
	Status   string     `json:"status"`
	Sequence int64      `json:"sequence"`
	Balances []*Balance `json:"balances"`
}

func (acct *Account) GetAddress() string {
	return acct.Address
}

func (acct *Account) GetLabel() string {
	return acct.Label
}

func (acct *Account) GetMemo() string {
	return acct.Memo
}

func (acct *Account) GetStatus() string {
	return acct.Status
}

func (acct *Account) SetStatus(status string) {
	acct.Status = status
}

func (acct *Account) GetSequence() int64 {
	return acct.Sequence
}

func (acct *Account) SetSequence(sequence int64) {
	acct.Sequence = sequence
}

func (acct *Account) GetSBalances() []*Balance {
	return acct.Balances
}
