package deposit

import (
	"github.com/xlalon/golee/pkg/ecode"
)

type IncomeCursor struct {
	chainCode string
	height    int64
	AccountIncomeCursor
}

type AccountIncomeCursor struct {
	address   string
	label     string
	txHash    string
	direction string
	index     int64
}

func newAccountIncomeCursor(address, label, txHash, direction string, index int64) *AccountIncomeCursor {
	aic := &AccountIncomeCursor{}
	if err := aic.SetAddress(address); err != nil {
		return nil
	}
	if err := aic.SetLabel(label); err != nil {
		return nil
	}
	if err := aic.SetTxHash(txHash); err != nil {
		return nil
	}
	if err := aic.SetDirection(direction); err != nil {
		return nil
	}
	if err := aic.SetIndex(index); err != nil {
		return nil
	}
	return aic
}

func NewIncomeCursor(chain string, height int64, address, label, txHash, direction string, index int64) *IncomeCursor {
	cursor := &IncomeCursor{}
	if err := cursor.SetChainCode(chain); err != nil {
		return nil
	}
	if err := cursor.SetHeight(height); err != nil {
		return nil
	}
	if err := cursor.SetTxHash(txHash); err != nil {
		return nil
	}
	if address != "" || label != "" || txHash != "" || direction != "" || index > 0 {
		if err := cursor.SetAccountIncomeCursor(newAccountIncomeCursor(address, label, txHash, direction, index)); err != nil {
			return nil
		}
	}
	return cursor
}

func (ic *IncomeCursor) ChainCode() string {
	return ic.chainCode
}

func (ic *IncomeCursor) SetChainCode(chainCode string) error {
	if ic.chainCode != "" {
		return ecode.CursorChange
	}
	if chainCode == "" {
		return ecode.CursorInvalid
	}
	ic.chainCode = chainCode
	return nil
}

func (ic *IncomeCursor) Height() int64 {
	return ic.height
}

func (ic *IncomeCursor) SetHeight(height int64) error {
	if height < 0 {
		return ecode.CursorInvalid
	}
	ic.height = height
	return nil
}

func (ic *IncomeCursor) TxHash() string {
	return ic.txHash
}

func (ic *IncomeCursor) SetTxHash(txHash string) error {
	ic.txHash = txHash
	return nil
}

func (ic *IncomeCursor) GetAccountIncomeCursor() AccountIncomeCursor {
	return ic.AccountIncomeCursor
}

func (ic *IncomeCursor) SetAccountIncomeCursor(ai *AccountIncomeCursor) error {
	ic.AccountIncomeCursor = *ai
	return nil
}

func (aic *AccountIncomeCursor) Address() string {
	return aic.address
}

func (aic *AccountIncomeCursor) SetAddress(address string) error {
	aic.address = address
	return nil
}

func (aic *AccountIncomeCursor) Label() string {
	return aic.label
}

func (aic *AccountIncomeCursor) SetLabel(label string) error {
	aic.label = label
	return nil
}

func (aic *AccountIncomeCursor) TxHash() string {
	return aic.txHash
}

func (aic *AccountIncomeCursor) SetTxHash(txHash string) error {
	aic.txHash = txHash
	return nil
}

func (aic *AccountIncomeCursor) Direction() string {
	return aic.direction
}

func (aic *AccountIncomeCursor) SetDirection(direction string) error {
	aic.direction = direction
	return nil
}

func (aic *AccountIncomeCursor) Index() int64 {
	return aic.index
}

func (aic *AccountIncomeCursor) SetIndex(index int64) error {
	if index < 0 {
		return ecode.ParameterInvalidError
	}
	aic.index = index
	return nil
}
