package deposit

import (
	"github.com/xlalon/golee/pkg/ecode"
)

type IncomeCursor struct {
	chainCode string

	height int64
	txHash string
	AccountIncomeCursor
}

type AccountIncomeCursor struct {
	address string
	label   string
	index   int64
}

func newAccountIncomeCursor(address, label string, index int64) *AccountIncomeCursor {
	aic := &AccountIncomeCursor{}
	if err := aic.SetAddress(address); err != nil {
		return nil
	}
	if err := aic.SetLabel(label); err != nil {
		return nil
	}
	if err := aic.SetIndex(index); err != nil {
		return nil
	}
	return aic
}

func NewIncomeCursor(cursorDTO *IncomeCursorDTO) *IncomeCursor {
	cursor := &IncomeCursor{}
	if err := cursor.SetChainCode(cursorDTO.ChainCode); err != nil {
		return nil
	}
	if err := cursor.SetHeight(cursorDTO.Height); err != nil {
		return nil
	}
	if err := cursor.SetTxHash(cursorDTO.TxHash); err != nil {
		return nil
	}
	if cursorDTO.Address != "" || cursorDTO.Label != "" || cursorDTO.Index > 0 {
		if err := cursor.SetAccountIncomeCursor(cursorDTO.Address, cursorDTO.Label, cursorDTO.Index); err != nil {
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
		return ecode.ParameterChangeError
	}
	if chainCode == "" {
		return ecode.ParameterInvalidError
	}
	ic.chainCode = chainCode
	return nil
}

func (ic *IncomeCursor) Height() int64 {
	return ic.height
}

func (ic *IncomeCursor) SetHeight(height int64) error {
	if height < 0 {
		return ecode.ParameterInvalidError
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

func (ic *IncomeCursor) SetAccountIncomeCursor(address, label string, index int64) error {
	if aic := newAccountIncomeCursor(address, label, index); aic != nil {
		ic.AccountIncomeCursor = *aic
	}
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

func (ic *IncomeCursor) ToCursorDTO() *IncomeCursorDTO {
	return &IncomeCursorDTO{
		ChainCode: ic.ChainCode(),
		Height:    ic.Height(),
		TxHash:    ic.TxHash(),
		Address:   ic.Address(),
		Label:     ic.Label(),
		Index:     ic.Index(),
	}
}
