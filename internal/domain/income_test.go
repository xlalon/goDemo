package domain

import (
	"fmt"
	"testing"

	"github.com/xlalon/golee/internal/infra/repository"
	rchainasset "github.com/xlalon/golee/internal/infra/repository/chainasset"
	rdeposit "github.com/xlalon/golee/internal/infra/repository/deposit"
	rwallet "github.com/xlalon/golee/internal/infra/repository/wallet"
	onchainConf "github.com/xlalon/golee/internal/onchain/conf"
	"github.com/xlalon/golee/internal/onchain/x"
	"github.com/xlalon/golee/pkg/database/mysql"
	"github.com/xlalon/golee/pkg/database/redis"
	"github.com/xlalon/golee/pkg/json"
)

var (
	_ = x.Init(&onchainConf.Config{
		Band: &onchainConf.ChainConfig{
			NodeUrl:           "https://api-bandchain-ia.cosmosia.notional.ventures/",
			ExternalNodeUrl:   "https://api-bandchain-ia.cosmosia.notional.ventures/",
			BlockTime:         7,
			IrreversibleBlock: 10,
			WalletDepositUrl:  "",
			WalletHotUrl:      "",
			SupportMemo:       true,
			DepositAddress:    "band1dkl8wga94803qygwdspwa5kxdfyjpt8zr0uzh9",
			HotAddress:        "band14fzetz7kn6wkuu0s0wv0lfmu795pl0wy3hesvu",
		},
		Waxp: &onchainConf.ChainConfig{
			NodeUrl:           "https://wax.greymass.com",
			ExternalNodeUrl:   "https://wax.greymass.com",
			BlockTime:         1,
			IrreversibleBlock: 100,
			WalletDepositUrl:  "",
			WalletHotUrl:      "",
			SupportMemo:       true,
			DepositAddress:    "",
			HotAddress:        "",
		},
	})

	testMysqlConf = &mysql.Config{DNS: "root:Xiao0000@tcp(127.0.0.1:3306)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"}
	testRedisConf = &redis.Config{
		Address:  "127.0.0.1",
		Port:     6379,
		Password: "",
		DB:       0,
	}

	testRepository = repository.NewRegistry(&repository.Config{
		Chain: &rchainasset.Config{
			Mysql: testMysqlConf,
			Redis: testRedisConf,
		},
		Deposit: &rdeposit.Config{
			Mysql: testMysqlConf,
			Redis: testRedisConf,
		},
		Wallet: &rwallet.Config{
			Mysql: testMysqlConf,
			Redis: testRedisConf,
		},
	})
	testIncome = NewIncome(
		testRepository.ChainRepository(),
		testRepository.DepositRepository(),
		testRepository.WalletRepository(),
	)
)

func TestService_ScanDeposits(t *testing.T) {
	deps, err := testIncome.scanDeposits("BAND", int64(15031628))
	fmt.Println("error:", err)
	for _, txn := range deps {
		fmt.Println("band deps:", fmt.Sprintf(
			"Id %d\n"+
				"Chain %s\n"+
				"Asset %s\n"+
				"TxHash %s\n"+
				"Sender %s\n"+
				"Receiver %s\n"+
				"AmountRaw %s\n"+
				"Precession %d\n"+
				"Amount %s\n"+
				"VOut %d\n"+
				"Status %s\n",
			txn.Id(),
			txn.Chain(),
			txn.Asset(),
			txn.TxHash(),
			txn.Sender(),
			txn.Receiver(),
			txn.AmountRaw(),
			txn.Precession(),
			txn.Amount(),
			txn.VOut(),
			txn.Status()))
	}
	testIncome.ScanDeposits("BAND", int64(15031628))
}

func TestIncome_GetCursor(t *testing.T) {
	json.PPrint("cursor", testIncome.GetCursor("BAND").ToCursorDTO())
}

//func TestIncome_SaveCursor(t *testing.T) {
//	fmt.Println(testIncome.SaveCursor(domain.NewIncomeCursor(&domain.IncomeCursorDTO{
//		ChainCode: "BAND",
//		Height:    15031627,
//		TxHash:    "8A4FF19A559EE8531D65D9AD99E7E333BF183B6AF273D22FD05E8F7A930192A0",
//		Address:   "band1dkl8wga94803qygwdspwa5kxdfyjpt8zr0uzh9",
//		Label:     "DEPOSIT",
//		Index:     1,
//	})))
//}
