package deposit

import (
	"github.com/xlalon/golee/common/math/decimal"
	"github.com/xlalon/golee/core/model/asset"
	"github.com/xlalon/golee/core/model/chain"
)

type DepositDto struct {
	Id int64

	Chain *chain.Chain

	TxHash string
	VOut   int64

	Senders []string

	Recipient string
	Memo      string

	Asset  *asset.Asset
	Amount decimal.Decimal

	Timestamp int64
	Height    int64
	Comment   interface{}

	Status Status
}
