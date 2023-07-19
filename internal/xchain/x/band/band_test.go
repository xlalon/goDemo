package band

import (
	"context"
	"fmt"
	"testing"

	"github.com/xlalon/golee/internal/xchain"
	"github.com/xlalon/golee/internal/xchain/conf"
	"github.com/xlalon/golee/pkg/json"
)

var (
	testBandChainConf = &conf.ChainConfig{
		NodeUrl:         "https://laozi1.bandchain.org/api",
		ExternalNodeUrl: "https://laozi1.bandchain.org/api",

		WalletDepositUrl: "",
		WalletHotUrl:     "",

		BlockTime:         7,
		IrreversibleBlock: 30,

		SupportMemo:    true,
		DepositAddress: "",
		HotAddress:     "",
	}
	testBand = New(testBandChainConf)
)

func TestBand_GetLatestHeight(t *testing.T) {
	height, err := testBand.GetLatestHeight(context.Background())
	if err != nil {
		t.Fatal("Band_GetLatestHeight error:", err)
	}
	fmt.Println("band height:", height)
}

func TestBand_GetTransfersByHash(t *testing.T) {
	txHash := "5202ECD08CC7938F5E1FA79BFFA50FAECC8D0F63335FDE0039C0296C6872C91D"
	tx, err := testBand.GetTransfersByHash(context.Background(), txHash)
	if err != nil {
		t.Fatal("Band_GetTxnByHash error:", err)
	}
	json.JPrint("band tx info", tx)
}

func TestBand_ScanTransfers(t *testing.T) {
	cursor := xchain.NewCursor(
		"BAND", 19230609, string(xchain.WalletLabelDeposit), testBandChainConf.DepositAddress, "", 0)
	txs, err := testBand.ScanTransfers(context.Background(), cursor)
	if err != nil {
		t.Fatal("Band_ScanTxn error:", err)
	}
	json.JPrint("band txs", txs)
}

func TestBand_NewAccount(t *testing.T) {
	ctx := context.Background()
	acctDeposit, err := testBand.NewAccount(ctx, xchain.WalletLabelDeposit)
	if err != nil {
		t.Fatal("Band_NewAccount error:", err)
	}
	json.JPrint("band new deposit account", acctDeposit)

	acctHot, err := testBand.NewAccount(ctx, xchain.WalletLabelHot)
	if err != nil {
		t.Fatal("Band_NewAccount error:", err)
	}
	json.JPrint("band new hot account", acctHot)
}

func TestBand_GetAccountBalance(t *testing.T) {
	balance, err := testBand.GetAccountBalance(context.Background(), xchain.Address(testBandChainConf.DepositAddress), "uband")
	if err != nil {
		t.Fatal("Band_GetBalance error:", err)
	}
	json.JPrint("band get balance", balance)
}

func TestBand_Transfer(t *testing.T) {
	fmt.Println("transfer...")
}
