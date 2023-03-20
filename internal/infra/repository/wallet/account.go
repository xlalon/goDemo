package wallet

import (
	"github.com/xlalon/golee/internal/service/wallet/domain"
	"time"

	"github.com/xlalon/golee/pkg/database/mysql"
)

func (d *Dao) Save(acct *domain.Account) error {
	var createAt = time.Time{}
	var version int64
	acctDB, err := d.getAccountById(acct.GetId())
	if err == nil && acctDB != nil {
		createAt = acctDB.CreatedAt
		version = acctDB.Version + 1
	}
	d.db.Model(&Account{Model: mysql.Model{ID: acct.GetId()}}).Save(&Account{
		Model: mysql.Model{
			ID:        acct.GetId(),
			CreatedAt: createAt,
		},
		Chain:   acct.GetChain(),
		Address: acct.GetAddress(),
		Label:   acct.GetLabel(),
		Memo:    acct.GetMemo(),
		Status:  acct.GetStatus(),
		Version: version,
	})
	return nil
}

func (d *Dao) GetAccountById(accountId int64) (*domain.Account, error) {
	acctDB, err := d.getAccountById(accountId)
	if err != nil {
		return nil, err
	}
	return d.acctDbToDomain(acctDB), nil
}

func (d *Dao) getAccountById(accountId int64) (*Account, error) {
	acctDB := &Account{}
	if err := d.db.Last(acctDB, "id = ?", accountId).Error; err != nil {
		return nil, err
	}
	return acctDB, nil
}

func (d *Dao) GetAccountByChainAddress(chain, address string) (*domain.Account, error) {
	acctDB := &Account{}
	if err := d.db.Last(acctDB, "chain = ? AND address = ?", chain, address).Error; err != nil {
		return nil, err
	}
	return d.acctDbToDomain(acctDB), nil
}

func (d *Dao) GetAccountByChainAddressMemo(chain, address, memo string) (*domain.Account, error) {
	acctDB := &Account{}
	if err := d.db.Last(acctDB, "chain = ? AND address = ? AND memo = ?", chain, address, memo).Error; err != nil {
		return nil, err
	}

	return d.acctDbToDomain(acctDB), nil
}

func (d *Dao) GetAccountsByChain(chain string) ([]*domain.Account, error) {
	var accountsDB []Account
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
	var accountsDB []Account
	if err := d.db.Find(&accountsDB, "chain = ? AND address in ?", chain, addresses).Error; err != nil {
		return nil, err
	}
	var accounts []*domain.Account
	for _, a := range accountsDB {
		accounts = append(accounts, d.acctDbToDomain(&a))
	}
	return accounts, nil
}

func (d *Dao) acctDbToDomain(acct *Account) *domain.Account {
	return domain.AccountFactory(&domain.AccountDTO{
		Id:      acct.ID,
		Chain:   acct.Chain,
		Address: acct.Address,
		Label:   acct.Label,
		Memo:    acct.Memo,
		Status:  acct.Status,
	})
}
