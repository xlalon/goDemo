package xchain

import (
	"github.com/xlalon/golee/pkg/math/decimal"
)

type TxnStatus string

const (
	TxnFailed  TxnStatus = "FAILED"
	TxnPending TxnStatus = "PENDING"
	TxnSuccess TxnStatus = "SUCCESS"
)

type Transfer struct {
	Chain         Chain       `json:"chain"`
	TxHash        string      `json:"tx_hash"`
	VOut          int64       `json:"v_out"`
	Sender        string      `json:"sender"`
	Recipient     *AccountDTO `json:"recipient"`
	CoinValue     Coin        `json:"coin"`
	Status        TxnStatus   `json:"status"`
	Confirmations int64       `json:"confirmations"`
	Height        int64       `json:"height"`
	Timestamp     int64       `json:"timestamp"`
	Comment       string      `json:"comment"`
}

type Fee struct {
	Chain    Chain           `json:"chain"`
	Coin     Coin            `json:"coin"`
	GasPrice decimal.Decimal `json:"gas_price"`
	GasLimit int64           `json:"gas_limit"`
}

type Receipt struct {
	TxHash string    `json:"tx_hash"`
	Fee    *Fee      `json:"fee"`
	Status TxnStatus `json:"status"`
	ErrLog string    `json:"err_log"`
}

type TransferCommand struct {
	Sender    *Wallet                `json:"sender"`
	Recipient *AccountDTO            `json:"receiver"`
	CoinValue Coin                   `json:"coin_value"`
	Fee       *Fee                   `json:"fee"`
	Extra     map[string]interface{} `json:"extra"`
}
