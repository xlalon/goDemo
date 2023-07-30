package deposit

import (
	"github.com/xlalon/golee/internal/domain/model"
	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/pkg/ecode"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type Deposit struct {
	model.IdentifiedDomainObject

	// core core info
	core Detail

	// extra info
	sender    string
	height    int64
	timestamp int64
	comment   string

	status Status
}

func DepositFactory(depositDTO *DepositDTO) *Deposit {
	deposit := &Deposit{}
	if err := deposit.SetId(depositDTO.Id); err != nil {
		return nil
	}
	if err := deposit.setCoreInfo(NewDepositItem(depositDTO)); err != nil {
		return nil
	}
	if err := deposit.setSender(depositDTO.Sender); err != nil {
		return nil
	}
	if err := deposit.setHeight(depositDTO.Height); err != nil {
		return nil
	}
	if err := deposit.setTimestamp(depositDTO.Timestamp); err != nil {
		return nil
	}
	if err := deposit.setComment(depositDTO.Comment); err != nil {
		return nil
	}
	if err := deposit.setStatus(depositDTO.Status); err != nil {
		return nil
	}
	return deposit
}

func (d *Deposit) Chain() string {
	return d.core.Chain()
}

func (d *Deposit) TxHash() string {
	return d.core.TxHash()
}

func (d *Deposit) VOut() int64 {
	return d.core.VOut()
}

func (d *Deposit) setCoreInfo(depositCoreInfo *Detail) error {
	if depositCoreInfo == nil {
		return ecode.DepositItemInvalid
	}
	d.core = *depositCoreInfo
	return nil
}

func (d *Deposit) Receiver() string {
	return d.core.Receiver()
}

func (d *Deposit) Memo() string {
	return d.core.Memo()
}

func (d *Deposit) Asset() chainasset.AssetCode {
	return d.core.assetValue.Asset()
}

func (d *Deposit) Amount() decimal.Decimal {
	return d.core.assetValue.Amount()
}

func (d *Deposit) Sender() string {
	return d.sender
}

func (d *Deposit) setSender(sender string) error {
	d.sender = sender
	return nil
}

func (d *Deposit) Height() int64 {
	return d.height
}

func (d *Deposit) setHeight(height int64) error {
	d.height = height
	return nil
}

func (d *Deposit) Timestamp() int64 {
	return d.timestamp
}

func (d *Deposit) setTimestamp(timestamp int64) error {
	d.timestamp = timestamp
	return nil
}

func (d *Deposit) Comment() string {
	return d.comment
}

func (d *Deposit) setComment(comment string) error {
	d.comment = comment
	return nil
}

func (d *Deposit) Status() Status {
	return d.status
}

func (d *Deposit) setStatus(status Status) error {
	if status != DepositStatusPending && status != DepositStatusFinished && status != DepositStatusCancelled && status != DepositStatusSwapped {
		return ecode.DepositStatusInvalid
	}
	d.status = status
	return nil
}

func (d *Deposit) ToDepositDTO() *DepositDTO {
	return &DepositDTO{
		Id: d.Id(),

		Chain:  d.Chain(),
		TxHash: d.TxHash(),
		VOut:   d.VOut(),

		Receiver: d.Receiver(),
		Memo:     d.Memo(),

		Asset:  string(d.Asset()),
		Amount: d.Amount(),

		Sender:    d.Sender(),
		Height:    d.Height(),
		Timestamp: d.Timestamp(),
		Comment:   d.Comment(),

		Status: d.Status(),
	}
}

type DepositDTO struct {
	Id int64 `json:"id"`

	Chain  string `json:"chain"`
	TxHash string `json:"tx_hash"`
	VOut   int64  `json:"v_out"`

	Asset  string          `json:"asset"`
	Amount decimal.Decimal `json:"amount"`

	Receiver string `json:"receiver"`
	Memo     string `json:"memo"`

	Sender    string `json:"sender"`
	Height    int64  `json:"height"`
	Status    Status `json:"status"`
	Timestamp int64  `json:"timestamp"`
	Comment   string `json:"comment"`
}
