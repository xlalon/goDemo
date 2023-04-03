package domain

import (
	"os"
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
	testMysqlConf = &mysql.Config{DSN: "root:Xiao0000@tcp(127.0.0.1:3306)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"}
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

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	x.Init(&onchainConf.Config{
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
	os.Exit(m.Run())
}

func TestService_ScanDeposits(t *testing.T) {
	//cursor := deposit.NewIncomeCursor(&deposit.IncomeCursorDTO{
	//	ChainCode: "BAND",
	//	Height:    15031627,
	//	Address:   "band1dkl8wga94803qygwdspwa5kxdfyjpt8zr0uzh9",
	//	Label:     "DEPOSIT",
	//	TxHash:    "8A4FF19A559EE8531D65D9AD99E7E333BF183B6AF273D22FD05E8F7A930192A0",
	//	Direction: "ASC",
	//	Index:     0,
	//})
	//onChainCursor := &onchain.Cursor{
	//	Chain:  onchain.Code(cursor.ChainCode()),
	//	Height: cursor.Height(),
	//	Account: &onchain.Account{
	//		Chain:   onchain.Code(cursor.ChainCode()),
	//		Address: cursor.Address(),
	//		Label:   onchain.Label(cursor.Label()),
	//	},
	//	TxHash:    cursor.TxHash(),
	//	Direction: onchain.Direction(cursor.Direction()),
	//	Index:     cursor.Index(),
	//}
	//	deps, err := testIncome.scanDeposits(onChainCursor)
	//	fmt.Println("error:", err)
	//	for _, txn := range deps {
	//		fmt.Println("band deps:", fmt.Sprintf(
	//			"Id %d\n"+
	//				"Chain %s\n"+
	//				"TxHash %s\n"+
	//				"VOut %d\n"+
	//				"Receiver %s\n"+
	//				"Memo %s\n"+
	//				"Asset %s\n"+
	//				"Amount %s\n"+
	//				"Sender %s\n"+
	//				"Height %d\n"+
	//				"Comment %v\n"+
	//				"Status %s\n",
	//			txn.Id(),
	//			txn.Chain(),
	//			txn.TxHash(),
	//			txn.VOut(),
	//			txn.Receiver(),
	//			txn.Memo(),
	//			txn.Asset(),
	//			txn.Amount(),
	//			txn.Sender(),
	//			txn.Height(),
	//			txn.Comment(),
	//			txn.Status()))
	//	}

	testIncome.ScanDeposits("BAND")
}

func TestIncome_GetCursor(t *testing.T) {
	json.PPrint("cursor", testIncome.incomeCursor("BAND"))
}

//func TestIncome_SaveCursor(t *testing.T) {
//	cursor := deposit.NewIncomeCursor(
//		"BAND",
//		15031626,
//		"band1dkl8wga94803qygwdspwa5kxdfyjpt8zr0uzh9",
//		"DEPOSIT",
//		"",
//		"ASC",
//		0)
//	onChainCursor := onchain.NewCursor(
//		onchain.Code(cursor.ChainCode()),
//		cursor.Height(),
//		cursor.Address(),
//		onchain.Label(cursor.Label()),
//		cursor.TxHash(),
//		onchain.Direction(cursor.Direction()),
//		cursor.Index(),
//	)
//	fmt.Println(testIncome.saveIncomeCursor(onChainCursor))
//}
