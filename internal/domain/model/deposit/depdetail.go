package deposit

import (
	"github.com/xlalon/golee/internal/domain/model/chainasset"
	"github.com/xlalon/golee/internal/domain/model/wallet"
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

type Detail struct {
	txnId      TxnId
	receiver   Receiver
	assetValue wallet.AssetValue
}

func NewDepositItem(depositDTO *DepositDTO) *Detail {
	item := &Detail{}
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

func (i *Detail) Chain() string {
	return i.txnId.chain
}

func (i *Detail) TxHash() string {
	return i.txnId.txHash
}

func (i *Detail) VOut() int64 {
	return i.txnId.vOut
}

func (i *Detail) setTxnId(chain, txHash string, vOut int64) error {
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

func (i *Detail) Receiver() string {
	return i.receiver.address
}

func (i *Detail) Memo() string {
	return i.receiver.memo
}

func (i *Detail) setReceiver(address, memo string) error {
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

func (i *Detail) Asset() chainasset.AssetCode {
	return i.assetValue.Asset()
}

func (i *Detail) Amount() decimal.Decimal {
	return i.assetValue.Amount()
}

func (i *Detail) setCoinValue(asset string, amount decimal.Decimal) error {
	if i.Asset() != "" || i.Amount().GreaterThanZero() {
		return ecode.DepositCoinValueChange
	}
	coinValue := wallet.NewAssetValue(asset, amount)
	if coinValue == nil {
		return ecode.DepositCoinValueInvalid
	}
	i.assetValue = *coinValue
	return nil
}
