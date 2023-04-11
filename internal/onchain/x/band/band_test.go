package band

import (
	"context"
	"fmt"
	"testing"

	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/internal/onchain/conf"
	"github.com/xlalon/golee/pkg/json"
)

var (
	testBandChainConf = &conf.ChainConfig{
		NodeUrl:         "https://band-api.ibs.team/",
		ExternalNodeUrl: "https://band-api.ibs.team/",

		WalletDepositUrl: "",
		WalletHotUrl:     "",

		BlockTime:         7,
		IrreversibleBlock: 30,

		SupportMemo:    true,
		DepositAddress: "band14fzetz7kn6wkuu0s0wv0lfmu795pl0wy3hesvu",
		HotAddress:     "band14fzetz7kn6wkuu0s0wv0lfmu795pl0wy3hesvu",
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

func TestBand_GetTxnByHash(t *testing.T) {
	tx, err := testBand.GetTxnByHash(context.Background(), "AE4CBD3ED9BD7A7CFDF532D8C241194CB15407D93DB0C35FAFA26A0DF3795AC7")
	if err != nil {
		t.Fatal("Band_GetTxnByHash error:", err)
	}
	json.PPrint("band tx info", tx)
}

func TestBand_ScanTxn(t *testing.T) {
	cursor := onchain.NewCursor(
		"BAND", 15663667, testBandChainConf.DepositAddress, onchain.AccountDeposit, "", onchain.CursorDirectionAsc, 0)
	txs, err := testBand.ScanTxn(context.Background(), cursor)
	if err != nil {
		t.Fatal("Band_ScanTxn error:", err)
	}
	json.PPrint("band txs", txs)
}

func TestBand_NewAccount(t *testing.T) {
	ctx := context.Background()
	acctDeposit, err := testBand.NewAccount(ctx, onchain.AccountDeposit)
	if err != nil {
		t.Fatal("Band_NewAccount error:", err)
	}
	json.PPrint("band new deposit account", acctDeposit)

	acctHot, err := testBand.NewAccount(ctx, onchain.AccountHot)
	if err != nil {
		t.Fatal("Band_NewAccount error:", err)
	}
	json.PPrint("band new hot account", acctHot)
}

func TestBand_GetAccount(t *testing.T) {
	acct, err := testBand.GetAccount(context.Background(), testBandChainConf.DepositAddress)
	if err != nil {
		t.Fatal("Band_GetAccount error:", err)
	}
	json.PPrint("band get account", acct)
}

func TestBand_Transfer(t *testing.T) {
	fmt.Println("transfer...")
}
