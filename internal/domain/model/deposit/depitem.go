package deposit

import (
	"github.com/xlalon/golee/pkg/ecode"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type DepositItem struct {
	chain    string
	asset    string
	txHash   string
	sender   string
	receiver string
	memo     string
	AmountVO
	vOut int64
}

func NewDepositItem(depositDTO *DepositDTO) *DepositItem {
	depositItem := &DepositItem{}
	if err := depositItem.setChain(depositDTO.Chain); err != nil {
		return nil
	}
	if err := depositItem.setAsset(depositDTO.Asset); err != nil {
		return nil
	}
	if err := depositItem.setTxHash(depositDTO.TxHash); err != nil {
		return nil
	}
	if err := depositItem.setSender(depositDTO.Sender); err != nil {
		return nil
	}
	if err := depositItem.setReceiver(depositDTO.Receiver); err != nil {
		return nil
	}
	if err := depositItem.setMemo(depositDTO.Memo); err != nil {
		return nil
	}
	if err := depositItem.setAmountInfo(depositDTO.Identity, depositDTO.AmountRaw, depositDTO.Precession, depositDTO.Amount); err != nil {
		return nil
	}
	if err := depositItem.setVOut(depositDTO.VOut); err != nil {
		return nil
	}
	return depositItem
}

func (dt *DepositItem) Chain() string {
	return dt.chain
}

func (dt *DepositItem) setChain(chain string) error {
	if dt.chain != "" {
		return ecode.ParameterChangeError
	}
	dt.chain = chain
	return nil
}

func (dt *DepositItem) Asset() string {
	return dt.asset
}

func (dt *DepositItem) setAsset(asset string) error {
	if dt.asset != "" {
		return ecode.ParameterChangeError
	}
	dt.asset = asset
	return nil
}

func (dt *DepositItem) TxHash() string {
	return dt.txHash
}

func (dt *DepositItem) setTxHash(txHash string) error {
	if dt.txHash != "" {
		return ecode.ParameterChangeError
	}
	dt.txHash = txHash
	return nil
}

func (dt *DepositItem) Sender() string {
	return dt.sender
}

func (dt *DepositItem) setSender(sender string) error {
	if dt.sender != "" {
		return ecode.ParameterChangeError
	}
	dt.sender = sender
	return nil
}

func (dt *DepositItem) Receiver() string {
	return dt.receiver
}

func (dt *DepositItem) setReceiver(receiver string) error {
	if dt.receiver != "" {
		return ecode.ParameterChangeError
	}
	dt.receiver = receiver
	return nil
}

func (dt *DepositItem) Memo() string {
	return dt.memo
}

func (dt *DepositItem) setMemo(memo string) error {
	if dt.memo != "" {
		return ecode.ParameterChangeError
	}
	dt.memo = memo
	return nil
}

func (dt *DepositItem) AmountInfo() AmountVO {
	return dt.AmountVO
}

func (dt *DepositItem) setAmountInfo(identity string, amountRaw decimal.Decimal, precession int64, amount decimal.Decimal) error {
	if dt.amount.GreaterThanZero() || dt.amountRaw.GreaterThanZero() || dt.identity != "" || dt.precession > 0 {
		return ecode.ParameterChangeError
	}
	dt.AmountVO = *NewAmountVO(identity, amountRaw, precession, amount)
	return nil
}

func (dt *DepositItem) VOut() int64 {
	return dt.vOut
}

func (dt *DepositItem) setVOut(vOut int64) error {
	if dt.vOut > 0 {
		return ecode.ParameterChangeError
	}
	dt.vOut = vOut
	return nil
}
