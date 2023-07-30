package wallet

import (
	"github.com/xlalon/golee/internal/domain/model"
	"github.com/xlalon/golee/pkg/ecode"
)

type Status string

const (
	AccountStatusValid   Status = "VALID"
	AccountStatusInValid        = "INVALID"
)

type Account struct {
	model.IdentifiedDomainObject

	chain string

	address string
	memo    string

	status string

	balances []*AssetValue
}

func NewAccount(acctDTO *AccountDTO) *Account {
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
	if err := acct.setMemo(acctDTO.Memo); err != nil {
		return nil
	}
	if err := acct.setStatus(acctDTO.Status); err != nil {
		return nil
	}
	var balances []*AssetValue
	for _, balanceDTO := range acctDTO.Balances {
		assetValue := NewAssetValue(balanceDTO.Asset, balanceDTO.Amount)
		if assetValue != nil {
			balances = append(balances, assetValue)
		}
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

func (acct *Account) Balances() []*AssetValue {
	return acct.balances
}

func (acct *Account) setBalances(balances []*AssetValue) error {
	acct.balances = balances
	return nil
}

func (acct *Account) ToAccountDTO() *AccountDTO {
	var balances []*BalanceDTO
	for _, balance := range acct.Balances() {
		balances = append(balances, &BalanceDTO{
			Amount: balance.Amount(),
			Asset:  string(balance.Asset()),
		})
	}
	return &AccountDTO{
		Id:       acct.Id(),
		Chain:    acct.Chain(),
		Address:  acct.Address(),
		Memo:     acct.Memo(),
		Status:   acct.Status(),
		Balances: balances,
	}
}
