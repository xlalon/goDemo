package domain

import (
	"github.com/xlalon/golee/pkg/math/decimal"
)

type Deposit struct {
	Id int64

	DepositItem
	Status Status
}

func NewDeposit(id int64, detail DepositItem, status Status) *Deposit {
	return &Deposit{
		Id:          id,
		DepositItem: detail,
		Status:      status,
	}
}

func (d *Deposit) GetId() int64 {
	return d.Id
}

func (d *Deposit) SetAsset(asset string) {
	d.DepositItem = *NewDepositItem(&DepositDTO{
		Chain:      d.chain,
		Asset:      asset,
		TxHash:     d.txHash,
		Sender:     d.sender,
		Receiver:   d.receiver,
		Memo:       d.memo,
		Identity:   d.identity,
		AmountRaw:  d.amountRaw,
		Precession: d.precession,
		Amount:     d.amount,
		VOut:       d.vOut,
		Status:     d.Status,
	})
}

func (d *Deposit) GetStatus() Status {
	return d.Status
}

func (d *Deposit) SetStatus(status Status) error {
	d.Status = status
	return nil
}

func (d *Deposit) SetPrecession(precession int64) {
	d.DepositItem = *NewDepositItem(&DepositDTO{
		Chain:      d.chain,
		Asset:      d.asset,
		TxHash:     d.txHash,
		Sender:     d.sender,
		Receiver:   d.receiver,
		Memo:       d.memo,
		Identity:   d.identity,
		AmountRaw:  d.amountRaw,
		Precession: precession,
		Amount:     d.amount,
		VOut:       d.vOut,
		Status:     d.Status,
	})
}

func (d *Deposit) SetAmount(amount decimal.Decimal) {
	d.DepositItem = *NewDepositItem(&DepositDTO{
		Chain:      d.chain,
		Asset:      d.asset,
		TxHash:     d.txHash,
		Sender:     d.sender,
		Receiver:   d.receiver,
		Memo:       d.memo,
		Identity:   d.identity,
		AmountRaw:  d.amountRaw,
		Precession: d.precession,
		Amount:     amount,
		VOut:       d.vOut,
		Status:     d.Status,
	})
}

func (d *Deposit) SetAmountRaw(amountRaw decimal.Decimal) {
	d.DepositItem = *NewDepositItem(&DepositDTO{
		Chain:      d.chain,
		Asset:      d.asset,
		TxHash:     d.txHash,
		Sender:     d.sender,
		Receiver:   d.receiver,
		Memo:       d.memo,
		Identity:   d.identity,
		AmountRaw:  amountRaw,
		Precession: d.precession,
		Amount:     d.amount,
		VOut:       d.vOut,
		Status:     d.Status,
	})
}

func (d *Deposit) IsValid() bool {
	return true
}

func (d *Deposit) ToDepositDTO() *DepositDTO {
	return &DepositDTO{
		Chain:      d.GetChain(),
		Asset:      d.GetAsset(),
		TxHash:     d.GetTxHash(),
		Sender:     d.GetSender(),
		Receiver:   d.GetReceiver(),
		Memo:       d.GetMemo(),
		Identity:   d.GetIdentity(),
		AmountRaw:  d.GetAmountRaw(),
		Precession: d.GetPrecession(),
		Amount:     d.GetAmount(),
		VOut:       d.GetVOut(),
		Status:     d.GetStatus(),
	}
}
