package waxp

import (
	"context"
	"fmt"
	"testing"

	"github.com/xlalon/golee/internal/xchain"
	"github.com/xlalon/golee/internal/xchain/conf"
	"github.com/xlalon/golee/pkg/json"
)

var (
	testWaxpChainConf = &conf.ChainConfig{
		NodeUrl:         "https://wax.hivebp.io",
		ExternalNodeUrl: "https://wax.hivebp.io",

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

func TestWaxp_GetTransfersByHash(t *testing.T) {
	txHash := "380b3b87baf40577ad82fd675512fd1fb2199a5a3d51726c561727f4084735cc"
	tx, err := testWaxp.GetTransfersByHash(context.Background(), txHash)
	if err != nil {
		t.Fatal("Waxp_GetTxnByHash error:", err)
	}
	json.JPrint("waxp tx info", tx)
}

func TestWaxp_ScanTransfers(t *testing.T) {
	cursor := xchain.NewCursor(
		"WAX", 0, string(xchain.WalletLabelDeposit), testWaxpChainConf.DepositAddress, "", 0)
	txs, err := testWaxp.ScanTransfers(context.Background(), cursor)
	if err != nil {
		t.Fatal("Waxp_ScanTxn error:", err)
	}
	json.JPrint("waxp scanned txs", txs)
}

func TestWaxp_NewAccount(t *testing.T) {
	ctx := context.Background()
	acctDeposit, err := testWaxp.NewAccount(ctx, xchain.WalletLabelDeposit)
	if err != nil {
		t.Fatal("Waxp_NewAccount error:", err)
	}
	json.JPrint("waxp new deposit account", acctDeposit)

	acctHot, err := testWaxp.NewAccount(ctx, xchain.WalletLabelHot)
	if err != nil {
		t.Fatal("Waxp_NewAccount error:", err)
	}
	json.JPrint("waxp new hot account", acctHot)
}

func TestWaxp_GetAccountBalance(t *testing.T) {
	balance, err := testWaxp.GetAccountBalance(context.Background(), xchain.Address(testWaxpChainConf.DepositAddress), "WAX")
	if err != nil {
		t.Fatal("Waxp_GetBalance error:", err)
	}
	json.JPrint("waxp get balance", balance)
}

func TestWaxp_Transfer(t *testing.T) {
	fmt.Println("transfer...")
}
