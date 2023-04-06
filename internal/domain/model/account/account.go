package account

import (
	"github.com/xlalon/golee/internal/domain/model"
	"github.com/xlalon/golee/pkg/ecode"
)

type Account struct {
	model.IdentifiedDomainObject

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
	if err := acct.SetId(acctDTO.Id); err != nil {
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
	var balances []*Balance
	for _, balanceDTO := range acctDTO.Balances {
		balances = append(balances, &Balance{
			Asset:      balanceDTO.Asset,
			Identity:   balanceDTO.Identity,
			Precession: balanceDTO.Precession,
			Amount:     balanceDTO.Amount,
			AmountRaw:  balanceDTO.AmountRaw,
		})
	}
	if err := acct.setBalances(balances); err != nil {
		return nil
	}
	return acct
}

func (acct *Account) Chain() string {
	return acct.chain
}

func (acct *Account) setChain(chain string) error {
	if acct.Chain() != "" {
		return ecode.AccountChainChange
	}
	if chain == "" {
		return ecode.AccountChainInvalid
	}
	acct.chain = chain
	return nil
}

func (acct *Account) Address() string {
	return acct.address
}

func (acct *Account) setAddress(address string) error {
	if acct.Address() != "" {
		return ecode.AccountAddressChange
	}
	if address == "" {
		return ecode.AccountAddressNull
	}
	acct.address = address
	return nil
}

func (acct *Account) Label() string {
	return acct.label
}

func (acct *Account) setLabel(label string) error {
	if acct.Label() != "" {
		return ecode.AccountLabelChange
	}
	if label == "" {
		return ecode.AccountLabelNull
	}
	acct.label = label
	return nil
}

func (acct *Account) Memo() string {
	return acct.memo
}

func (acct *Account) setMemo(memo string) error {
	if acct.Memo() != "" {
		return ecode.AccountMemoChange
	}
	acct.memo = memo
	return nil
}

func (acct *Account) Status() string {
	return acct.status
}

func (acct *Account) setStatus(status string) error {
	if status == "" {
		return ecode.AccountStatusInvalid
	}
	acct.status = status
	return nil
}

func (acct *Account) Sequence() int64 {
	return acct.sequence
}

func (acct *Account) setSequence(sequence int64) error {
	if sequence < 0 {
		return ecode.AccountSequenceInvalid
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
	var balances []*BalanceDTO
	for _, balance := range acct.Balances() {
		balances = append(balances, balance.ToBalanceDTO())
	}
	return &AccountDTO{
		Id:       acct.Id(),
		Chain:    acct.Chain(),
		Address:  acct.Address(),
		Label:    acct.Label(),
		Memo:     acct.Memo(),
		Status:   acct.Status(),
		Sequence: acct.Sequence(),
		Balances: balances,
	}
}
