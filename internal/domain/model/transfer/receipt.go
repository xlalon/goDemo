package transfer

type TxnStatus string

type Receipt struct {
	TxHash string    `json:"tx_hash"`
	Status TxnStatus `json:"status"`
}
