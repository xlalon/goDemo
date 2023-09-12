package deposit

import (
	"context"
	"github.com/xlalon/golee/common/database/mysql"
	"github.com/xlalon/golee/common/math/rand"
	"github.com/xlalon/golee/core/model/account"
	"github.com/xlalon/golee/x"
	"github.com/xlalon/golee/x/chain"
)

func (s *Service) NewAddress(chainCode string) (*account.Account, error) {
	c, err := x.GetChain(chainCode)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	wallet, err := c.GetWallet(ctx, chain.Deposit)
	if err != nil {
		return nil, err
	}
	acct, err := wallet.NewAccount(ctx)
	if err != nil {
		return nil, err
	}
	address, err := acct.Address(ctx)
	if err != nil {
		return nil, err
	}
	return account.NewAccount(mysql.NextID(), address, rand.DigitalMemo(), string(account.StatusActive)), nil
}
