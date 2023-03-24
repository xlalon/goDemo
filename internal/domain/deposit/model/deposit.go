package model

import (
	"github.com/xlalon/golee/pkg/ecode"
)

type Deposit struct {
	id int64

	DepositItem

	status Status
}

func DepositFactory(depositDTO *DepositDTO) *Deposit {
	deposit := &Deposit{}
	if err := deposit.setId(depositDTO.Id); err != nil {
		return nil
	}
	if err := deposit.setDepositItem(NewDepositItem(depositDTO)); err != nil {
		return nil
	}
	if err := deposit.setStatus(depositDTO.Status); err != nil {
		return nil
	}
	return deposit
}

func (d *Deposit) Id() int64 {
	return d.id
}

func (d *Deposit) setId(id int64) error {
	if d.Id() != 0 {
		return ecode.ParameterChangeError
	}
	if id <= 0 {
		return ecode.ParameterInvalidError
	}
	d.id = id
	return nil
}

func (d *Deposit) setDepositItem(depositItem *DepositItem) error {
	if depositItem == nil {
		return ecode.ParameterNullError
	}
	d.DepositItem = *depositItem
	return nil
}

func (d *Deposit) Status() Status {
	return d.status
}

func (d *Deposit) setStatus(status Status) error {
	if status != DepositStatusPending && status != DepositStatusFinished && status != DepositStatusCancelled && status != DepositStatusSwapped {
		return ecode.ParameterInvalidError
	}
	d.status = status
	return nil
}

func (d *Deposit) ToDepositDTO() *DepositDTO {
	return &DepositDTO{
		Id:         d.Id(),
		Chain:      d.Chain(),
		Asset:      d.Asset(),
		TxHash:     d.TxHash(),
		Sender:     d.Sender(),
		Receiver:   d.Receiver(),
		Memo:       d.Memo(),
		Identity:   d.Identity(),
		AmountRaw:  d.AmountRaw(),
		Precession: d.Precession(),
		Amount:     d.Amount(),
		VOut:       d.VOut(),
		Status:     d.Status(),
	}
}
