package domain

type Account struct {
	Id int64

	chain    string
	address  string
	label    string
	memo     string
	status   string
	sequence int64
	balances []*Balance
}

func AccountFactory(dto *AccountDTO) *Account {
	return &Account{
		Id:       dto.Id,
		chain:    dto.Chain,
		address:  dto.Address,
		label:    dto.Label,
		memo:     dto.Memo,
		status:   dto.Status,
		sequence: dto.Sequence,
		balances: dto.Balances,
	}
}

func (acct *Account) GetId() int64 {
	return acct.Id
}

func (acct *Account) GetChain() string {
	return acct.chain
}

func (acct *Account) GetAddress() string {
	return acct.address
}

func (acct *Account) GetLabel() string {
	return acct.label
}

func (acct *Account) GetMemo() string {
	return acct.memo
}

func (acct *Account) GetStatus() string {
	return acct.status
}

func (acct *Account) SetStatus(status string) {
	acct.status = status
}

func (acct *Account) GetSequence() int64 {
	return acct.sequence
}

func (acct *Account) SetSequence(sequence int64) {
	acct.sequence = sequence
}

func (acct *Account) GetSBalances() []*Balance {
	return acct.balances
}

func (acct *Account) ToAccountDTO() *AccountDTO {
	return &AccountDTO{
		Id:       acct.GetId(),
		Chain:    acct.GetChain(),
		Address:  acct.GetAddress(),
		Label:    acct.GetLabel(),
		Memo:     acct.GetMemo(),
		Status:   acct.GetStatus(),
		Sequence: acct.GetSequence(),
		Balances: acct.GetSBalances(),
	}
}
