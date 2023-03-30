package wallet

import "github.com/xlalon/golee/pkg/ecode"

type Account struct {
	id int64

	chain    string
	address  string
	label    string
	memo     string
	status   string
	sequence int64
	balances []*Balance
}

func AccountFactory(acctDTO *AccountDTO) *Account {
	acct := &Account{}
	if err := acct.setId(acctDTO.Id); err != nil {
		return nil
	}
	if err := acct.setChain(acctDTO.Chain); err != nil {
		return nil
	}
	if err := acct.setAddress(acctDTO.Address); err != nil {
		return nil
	}
	if err := acct.setLabel(acctDTO.Label); err != nil {
		return nil
	}
	if err := acct.setMemo(acctDTO.Memo); err != nil {
		return nil
	}
	if err := acct.setStatus(acctDTO.Status); err != nil {
		return nil
	}
	if err := acct.setSequence(acctDTO.Sequence); err != nil {
		return nil
	}
	if err := acct.setBalances(acctDTO.Balances); err != nil {
		return nil
	}
	return acct
}

func (acct *Account) Id() int64 {
	return acct.id
}

func (acct *Account) setId(id int64) error {
	if acct.Id() != 0 {
		return ecode.ParameterChangeError
	}
	if id <= 0 {
		return ecode.ParameterInvalidError
	}
	acct.id = id
	return nil
}

func (acct *Account) Chain() string {
	return acct.chain
}

func (acct *Account) setChain(chain string) error {
	if acct.Chain() != "" {
		return ecode.ParameterChangeError
	}
	if chain == "" {
		return ecode.ParameterNullError
	}
	acct.chain = chain
	return nil
}

func (acct *Account) Address() string {
	return acct.address
}

func (acct *Account) setAddress(address string) error {
	if acct.Address() != "" {
		return ecode.ParameterChangeError
	}
	if address == "" {
		return ecode.ParameterNullError
	}
	acct.address = address
	return nil
}

func (acct *Account) Label() string {
	return acct.label
}

func (acct *Account) setLabel(label string) error {
	if acct.Label() != "" {
		return ecode.ParameterChangeError
	}
	if label == "" {
		return ecode.ParameterNullError
	}
	acct.label = label
	return nil
}

func (acct *Account) Memo() string {
	return acct.memo
}

func (acct *Account) setMemo(memo string) error {
	if acct.Memo() != "" {
		return ecode.ParameterChangeError
	}
	if memo == "" {
		return ecode.ParameterNullError
	}
	acct.memo = memo
	return nil
}

func (acct *Account) Status() string {
	return acct.status
}

func (acct *Account) setStatus(status string) error {
	if status == "" {
		return ecode.ParameterNullError
	}
	acct.status = status
	return nil
}

func (acct *Account) Sequence() int64 {
	return acct.sequence
}

func (acct *Account) setSequence(sequence int64) error {
	if sequence < 0 {
		return ecode.ParameterInvalidError
	}
	acct.sequence = sequence
	return nil
}

func (acct *Account) Balances() []*Balance {
	return acct.balances
}

func (acct *Account) setBalances(balances []*Balance) error {
	acct.balances = balances
	return nil
}

func (acct *Account) ToAccountDTO() *AccountDTO {
	return &AccountDTO{
		Id:       acct.Id(),
		Chain:    acct.Chain(),
		Address:  acct.Address(),
		Label:    acct.Label(),
		Memo:     acct.Memo(),
		Status:   acct.Status(),
		Sequence: acct.Sequence(),
		Balances: acct.Balances(),
	}
}
