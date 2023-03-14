package dao

import (
	"github.com/xlalon/golee/internal/service/wallet/domain"
	"github.com/xlalon/golee/internal/service/wallet/repository/model"
)

func (d *Dao) NewAccount(acct *domain.Account) error {
	return d.db.Create(&model.Account{
		Chain:   acct.Chain,
		Address: acct.Address,
		Label:   acct.Label,
		Memo:    acct.Memo,
		Status:  acct.Status,
	}).Error
}

func (d *Dao) GetAccountByChainAddress(chain, address string) (*domain.Account, error) {
	acctDB := &model.Account{}
	if err := d.db.Last(acctDB, "chain = ? AND address = ?", chain, address).Error; err != nil {
		return nil, err
	}
	return d.acctDbToDomain(acctDB), nil
}

func (d *Dao) GetAccountByChainAddressMemo(chain, address, memo string) (*domain.Account, error) {
	acctDB := &model.Account{}
	if err := d.db.Last(acctDB, "chain = ? AND address = ? AND memo = ?", chain, address, memo).Error; err != nil {
		return nil, err
	}

	return d.acctDbToDomain(acctDB), nil
}

func (d *Dao) GetAccountsByChain(chain string) ([]*domain.Account, error) {
	var accountsDB []model.Account
	if err := d.db.Find(&accountsDB, "chain = ?", chain).Error; err != nil {
		return nil, err
	}
	var accounts []*domain.Account
	for _, a := range accountsDB {
		accounts = append(accounts, d.acctDbToDomain(&a))
	}
	return accounts, nil
}

func (d *Dao) GetAccountsByChainAddresses(chain string, addresses []string) ([]*domain.Account, error) {
	var accountsDB []model.Account
	if err := d.db.Find(&accountsDB, "chain = ? AND address in ?", chain, addresses).Error; err != nil {
		return nil, err
	}
	var accounts []*domain.Account
	for _, a := range accountsDB {
		accounts = append(accounts, d.acctDbToDomain(&a))
	}
	return accounts, nil
}

func (d *Dao) acctDbToDomain(acct *model.Account) *domain.Account {
	return &domain.Account{
		Chain:    acct.Chain,
		Address:  acct.Address,
		Label:    acct.Label,
		Memo:     acct.Memo,
		Status:   acct.Status,
		Sequence: 0,
		Balances: nil,
	}
}
