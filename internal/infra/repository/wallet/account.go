package wallet

import (
	"time"

	"github.com/xlalon/golee/internal/domain/model/wallet"
	"github.com/xlalon/golee/pkg/database/mysql"
)

func (d *Dao) Save(acct *wallet.Account) error {
	var createdAt time.Time
	var version int64
	acctDB, err := d.getAccountById(acct.Id())
	if err == nil && acctDB != nil {
		createdAt = acctDB.CreatedAt
		version = acctDB.Version + 1
	}
	d.db.Model(&Account{Model: mysql.Model{ID: acct.Id()}}).Save(&Account{
		Model: mysql.Model{
			ID:        acct.Id(),
			CreatedAt: createdAt,
		},
		Chain:   acct.Chain(),
		Address: acct.Address(),
		Label:   acct.Label(),
		Memo:    acct.Memo(),
		Status:  acct.Status(),
		Version: version,
	})
	return nil
}

func (d *Dao) GetAccountById(accountId int64) (*wallet.Account, error) {
	return d.acctDbToDomain(d.getAccountById(accountId))
}

func (d *Dao) GetAccountByChainAddress(chain, address string) (*wallet.Account, error) {
	return d.acctDbToDomain(d.getAccountByChainAddress(chain, address))
}

func (d *Dao) GetAccountByChainAddressMemo(chain, address, memo string) (*wallet.Account, error) {
	return d.acctDbToDomain(d.getAccountByChainAddressMemo(chain, address, memo))
}

func (d *Dao) GetAccountsByChain(chain string) ([]*wallet.Account, error) {
	return d.accountsDbToDomain(d.getAccountsByChain(chain))
}

func (d *Dao) GetAccountsByChainAddresses(chain string, addresses []string) ([]*wallet.Account, error) {
	return d.accountsDbToDomain(d.getAccountsByChainAddresses(chain, addresses))
}

func (d *Dao) getAccountById(accountId int64) (*Account, error) {
	acctDB := &Account{}
	if err := d.db.Last(acctDB, "id = ?", accountId).Error; err != nil {
		return nil, err
	}
	return acctDB, nil
}

func (d *Dao) getAccountByChainAddress(chain, address string) (*Account, error) {
	acctDB := &Account{}
	if err := d.db.Last(acctDB, "chain = ? AND address = ?", chain, address).Error; err != nil {
		return nil, err
	}
	return acctDB, nil
}

func (d *Dao) getAccountByChainAddressMemo(chain, address, memo string) (*Account, error) {
	acctDB := &Account{}
	if err := d.db.Last(acctDB, "chain = ? AND address = ? AND memo = ?", chain, address, memo).Error; err != nil {
		return nil, err
	}
	return acctDB, nil
}

func (d *Dao) getAccountsByChain(chain string) ([]Account, error) {
	var accountsDB []Account
	if err := d.db.Find(&accountsDB, "chain = ?", chain).Error; err != nil {
		return nil, err
	}
	return accountsDB, nil

}

func (d *Dao) getAccountsByChainAddresses(chain string, addresses []string) ([]Account, error) {
	var accountsDB []Account
	if err := d.db.Find(&accountsDB, "chain = ? AND address in ?", chain, addresses).Error; err != nil {
		return nil, err
	}
	return accountsDB, nil
}

func (d *Dao) accountsDbToDomain(accountsDB []Account, err error) ([]*wallet.Account, error) {
	if err != nil || accountsDB == nil || len(accountsDB) == 0 {
		return nil, err
	}
	var accountsDM []*wallet.Account
	for _, acctDB := range accountsDB {
		if acctDM, _ := d.acctDbToDomain(&acctDB, nil); acctDM != nil {
			accountsDM = append(accountsDM, acctDM)
		}
	}
	return accountsDM, nil
}

func (d *Dao) acctDbToDomain(acctDB *Account, err error) (*wallet.Account, error) {
	if err != nil || acctDB == nil {
		return nil, err
	}
	return wallet.AccountFactory(&wallet.AccountDTO{
		Id:      acctDB.ID,
		Chain:   acctDB.Chain,
		Address: acctDB.Address,
		Label:   acctDB.Label,
		Memo:    acctDB.Memo,
		Status:  acctDB.Status,
	}), nil
}
