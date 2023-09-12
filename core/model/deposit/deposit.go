package deposit

import (
	"github.com/xlalon/golee/common/math/decimal"
	"github.com/xlalon/golee/core/model/asset"
	"github.com/xlalon/golee/core/model/chain"
)

type Status string

const (
	StatusFailed  Status = "FAILED"
	StatusPending Status = "PENDING"
	StatusSuccess Status = "SUCCESS"
)

type Deposit struct {
	id int64

	chain *chain.Chain

	txHash string
	vOut   int64

	senders []string

	recipient string
	memo      string

	asset  *asset.Asset
	amount decimal.Decimal

	timestamp int64
	height    int64
	comment   interface{}

	status Status
}

func NewDeposit(dto *DepositDto) *Deposit {
	return &Deposit{
		id:        dto.Id,
		chain:     dto.Chain,
		txHash:    dto.TxHash,
		vOut:      dto.VOut,
		senders:   dto.Senders,
		recipient: dto.Recipient,
		memo:      dto.Memo,
		asset:     dto.Asset,
		amount:    dto.Amount,
		timestamp: dto.Timestamp,
		height:    dto.Height,
		comment:   dto.Comment,
		status:    dto.Status,
	}
}

func (d *Deposit) Id() int64 {
	return d.id
}
