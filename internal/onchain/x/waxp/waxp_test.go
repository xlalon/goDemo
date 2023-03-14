package waxp

import (
	"fmt"
	"github.com/xlalon/golee/internal/onchain"
	"testing"

	"github.com/xlalon/golee/internal/onchain/conf"
	"github.com/xlalon/golee/pkg/json"
)

var (
	testWaxpChainConf = &conf.ChainConfig{
		NodeUrl:         "https://wax.greymass.com",
		ExternalNodeUrl: "https://wax.greymass.com",

		WalletDepositUrl: "",
		WalletHotUrl:     "",

		BlockTime:         2,
		IrreversibleBlock: 30,

		SupportMemo:    true,
		DepositAddress: "",
		HotAddress:     "",
	}
	testWaxp = New(testWaxpChainConf)
)

func TestWaxp_GetLatestHeight(t *testing.T) {
	height, err := testWaxp.GetLatestHeight()
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("waxp height:", height)
}

func TestWaxp_GetTxnByHash(t *testing.T) {
	tx, err := testWaxp.GetTxnByHash("f3b8a14016b9fbcf6b0d834c78e880119304ff58c3c5876541c047a7180dd8b6")
	if err != nil {
		fmt.Println("error:", err)
	}
	json.PPrint("waxp tx info:", tx)
}

func TestWaxp_ScanTxn(t *testing.T) {
	acct := &onchain.Account{
		Chain:   "WAX",
		Address: "coinex111111",
	}
	txs, err := testWaxp.ScanTxn(acct)
	if err != nil {
		fmt.Println("error:", err)
	}
	json.PPrint("waxp scanned txs:", txs)
}

func TestWaxp_NewAccount(t *testing.T) {
	acctDeposit, err := testWaxp.NewAccount("deposit")
	if err != nil {
		fmt.Println("error:", err)
	}
	json.PPrint("waxp new deposit account", acctDeposit)

	acctHot, err := testWaxp.NewAccount("hot")
	if err != nil {
		fmt.Println("error:", err)
	}
	json.PPrint("waxp new hot account", acctHot)
}

func TestWaxp_GetAccount(t *testing.T) {
	acct, err := testWaxp.GetAccount("")
	if err != nil {
		fmt.Println("error:", err)
	}
	json.PPrint("waxp get account:", acct)
}

func TestWaxp_Transfer(t *testing.T) {
	fmt.Println("transfer...")
}
