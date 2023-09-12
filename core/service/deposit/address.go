package deposit

import (
	"context"
	"github.com/xlalon/golee/common/database/mysql"
	"github.com/xlalon/golee/common/math/rand"
	"github.com/xlalon/golee/core/model/account"
	"github.com/xlalon/golee/x"
	"github.com/xlalon/golee/x/chain"
)

func (s *Service) NewAddress(chainCode string) *account.Account {
	c, err := x.GetChain(chainCode)
	if err != nil {
		return nil
	}
	ctx := context.Background()
	wallet, err := c.GetWallet(ctx, chain.Deposit)
	if err != nil {
		return nil
	}
	acct, err := wallet.NewAccount(ctx)
	if err != nil {
		return nil
	}
	address, err := acct.Address(ctx)
	if err != nil {
		return nil
	}
	return account.NewAccount(mysql.NextID(), address, rand.DigitalMemo(), int64(account.StatusActive))
}
