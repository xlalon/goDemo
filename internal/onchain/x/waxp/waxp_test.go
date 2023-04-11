package waxp

import (
	"context"
	"fmt"
	"testing"

	"github.com/xlalon/golee/internal/onchain"
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
	height, err := testWaxp.GetLatestHeight(context.Background())
	if err != nil {
		t.Fatal("Waxp_GetLatestHeight error:", err)
	}
	fmt.Println("waxp height:", height)
}

func TestWaxp_GetTxnByHash(t *testing.T) {
	tx, err := testWaxp.GetTxnByHash(context.Background(), "f3b8a14016b9fbcf6b0d834c78e880119304ff58c3c5876541c047a7180dd8b6")
	if err != nil {
		t.Fatal("Waxp_GetTxnByHash error:", err)
	}
	json.PPrint("waxp tx info", tx)
}

func TestWaxp_ScanTxn(t *testing.T) {
	cursor := onchain.NewCursor(
		"WAX", 0, testWaxpChainConf.DepositAddress, onchain.AccountDeposit, "", onchain.CursorDirectionAsc, 0)
	txs, err := testWaxp.ScanTxn(context.Background(), cursor)
	if err != nil {
		t.Fatal("Waxp_ScanTxn error:", err)
	}
	json.PPrint("waxp scanned txs", txs)
}

func TestWaxp_NewAccount(t *testing.T) {
	ctx := context.Background()
	acctDeposit, err := testWaxp.NewAccount(ctx, onchain.AccountDeposit)
	if err != nil {
		t.Fatal("Waxp_NewAccount error:", err)
	}
	json.PPrint("waxp new deposit account", acctDeposit)

	acctHot, err := testWaxp.NewAccount(ctx, onchain.AccountHot)
	if err != nil {
		t.Fatal("Waxp_NewAccount error:", err)
	}
	json.PPrint("waxp new hot account", acctHot)
}

func TestWaxp_GetAccount(t *testing.T) {
	acct, err := testWaxp.GetAccount(context.Background(), testWaxpChainConf.DepositAddress)
	if err != nil {
		t.Fatal("Waxp_GetAccount error:", err)
	}
	json.PPrint("waxp get account", acct)
}

func TestWaxp_Transfer(t *testing.T) {
	fmt.Println("transfer...")
}
