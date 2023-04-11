package transfer

import (
	"github.com/xlalon/golee/internal/domain/model/account"
	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type Transfer struct {
	chain                        string
	fromAccount, receiverAccount *account.Account
	identity                     string
	amount                       decimal.Decimal
	extra                        map[string]interface{}

	repo       account.AccountRepository
	onchainSvc *onchain.Service
}

func (t *Transfer) Transfer() {
	//
}
