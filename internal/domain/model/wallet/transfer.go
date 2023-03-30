package wallet

import (
	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type Transfer struct {
	chain                        string
	fromAccount, receiverAccount *Account
	identity                     string
	amount                       decimal.Decimal
	extra                        interface{}

	repo       WalletRepository
	onchainSvc *onchain.Service
}

func (t *Transfer) Transfer() {
	//
}
