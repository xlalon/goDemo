package wallet

import (
	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/internal/service/wallet/domain"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type Transfer struct {
	chain                        string
	fromAccount, receiverAccount *domain.Account
	identity                     string
	amount                       decimal.Decimal
	extra                        interface{}

	repo       domain.WalletRepository
	onchainSvc *onchain.Service
}

func (t *Transfer) Transfer() {
	//
}
