package deposit

import (
	"github.com/xlalon/golee/pkg/ecode"
	"github.com/xlalon/golee/pkg/math/decimal"
)

type TxnId struct {
	chain  string
	txHash string
	vOut   int64
}

type Receiver struct {
	address string
	memo    string
}

type Item struct {
	txnId      TxnId
	receiver   Receiver
	assetValue AssetValue
}

func NewDepositItem(depositDTO *DepositDTO) *Item {
	item := &Item{}
	if err := item.setTxnId(depositDTO.Chain, depositDTO.TxHash, depositDTO.VOut); err != nil {
		return nil
	}
	if err := item.setReceiver(depositDTO.Receiver, depositDTO.Memo); err != nil {
		return nil
	}
	if err := item.setCoinValue(depositDTO.Asset, depositDTO.Amount); err != nil {
		return nil
	}
	return item
}

func (i *Item) Chain() string {
	return i.txnId.chain
}

func (i *Item) TxHash() string {
	return i.txnId.txHash
}

func (i *Item) VOut() int64 {
	return i.txnId.vOut
}

func (i *Item) setTxnId(chain, txHash string, vOut int64) error {
	if i.Chain() != "" || i.TxHash() != "" || i.VOut() > 0 {
		return ecode.DepositTxIdChange
	}
	if chain == "" {
		return ecode.DepositTxIdChainInvalid
	}
	if txHash == "" {
		return ecode.DepositTxIdHashInvalid
	}
	if vOut < 0 {
		return ecode.DepositTxIdVOutInvalid
	}
	i.txnId = TxnId{
		chain:  chain,
		txHash: txHash,
		vOut:   vOut,
	}
	return nil
}

func (i *Item) Receiver() string {
	return i.receiver.address
}

func (i *Item) Memo() string {
	return i.receiver.memo
}

func (i *Item) setReceiver(address, memo string) error {
	if i.Receiver() != "" || i.Memo() != "" {
		return ecode.DepositAssetInvalid
	}
	if address == "" {
		return ecode.DepositReceiverAddressInvalid
	}
	i.receiver = Receiver{
		address: address,
		memo:    memo,
	}
	return nil
}

func (i *Item) Asset() string {
	return i.assetValue.Asset()
}

func (i *Item) Amount() decimal.Decimal {
	return i.assetValue.Amount()
}

func (i *Item) setCoinValue(asset string, amount decimal.Decimal) error {
	if i.Asset() != "" || i.Amount().GreaterThanZero() {
		return ecode.DepositCoinValueChange
	}
	coinValue := NewAssetValue(asset, amount)
	if coinValue == nil {
		return ecode.DepositCoinValueInvalid
	}
	i.assetValue = *coinValue
	return nil
}
