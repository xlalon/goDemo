package domain

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

func NewDepositItem(dto *DepositDTO) *DepositItem {
	return &DepositItem{
		chain:    dto.Chain,
		asset:    dto.Asset,
		txHash:   dto.TxHash,
		sender:   dto.Sender,
		receiver: dto.Receiver,
		memo:     dto.Memo,
		AmountVO: NewAmountVO(dto.Identity, dto.AmountRaw, dto.Precession, dto.Amount),
		vOut:     dto.VOut,
	}
}

func (dt *DepositItem) GetChain() string {
	return dt.chain
}

func (dt *DepositItem) GetAsset() string {
	return dt.asset
}

func (dt *DepositItem) GetTxHash() string {
	return dt.txHash
}

func (dt *DepositItem) GetSender() string {
	return dt.sender
}

func (dt *DepositItem) GetReceiver() string {
	return dt.receiver
}

func (dt *DepositItem) GetMemo() string {
	return dt.memo
}

func (dt *DepositItem) GetAmountInfo() AmountVO {
	return dt.AmountVO
}

func (dt *DepositItem) GetVOut() int64 {
	return dt.vOut
}
