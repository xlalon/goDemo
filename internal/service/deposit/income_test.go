package deposit

import (
	"fmt"
	"testing"

	"github.com/xlalon/golee/internal/onchain"
	onchainConf "github.com/xlalon/golee/internal/onchain/conf"
	"github.com/xlalon/golee/internal/onchain/x"
	chainSvc "github.com/xlalon/golee/internal/service/chain"
	chainConf "github.com/xlalon/golee/internal/service/chain/conf"
	"github.com/xlalon/golee/internal/service/deposit/conf"
	"github.com/xlalon/golee/internal/service/wallet"
	walletConf "github.com/xlalon/golee/internal/service/wallet/conf"
	"github.com/xlalon/golee/pkg/database/mysql"
	"github.com/xlalon/golee/pkg/database/redis"
)

var (
	testOnchainConf = &onchainConf.Config{
		Band: &onchainConf.ChainConfig{
			NodeUrl:           "https://api-bandchain-ia.cosmosia.notional.ventures/",
			ExternalNodeUrl:   "https://api-bandchain-ia.cosmosia.notional.ventures/",
			BlockTime:         7,
			IrreversibleBlock: 10,
			WalletDepositUrl:  "",
			WalletHotUrl:      "",
			SupportMemo:       true,
			DepositAddress:    "band1ggq8us6lh4c8hr4624xnrlud6q5lqhklakysnd",
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
	}
	_ = x.Init(testOnchainConf)

	testIncomeConf = &conf.Config{
		Mysql: &mysql.Config{DNS: "mycat:p123456@tcp(127.0.0.1:8066)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"},
		Redis: &redis.Config{
			Address:  "127.0.0.1",
			Port:     6379,
			Password: "",
			DB:       0,
		},
	}
	testChainConf = &chainConf.Config{
		Mysql: &mysql.Config{DNS: "mycat:p123456@tcp(127.0.0.1:8066)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"},
		Redis: &redis.Config{
			Address:  "127.0.0.1",
			Port:     6379,
			Password: "",
			DB:       0,
		},
	}
	testWalletConf = &walletConf.Config{
		Mysql: &mysql.Config{DNS: "mycat:p123456@tcp(127.0.0.1:8066)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"},
		Redis: &redis.Config{
			Address:  "127.0.0.1",
			Port:     6379,
			Password: "",
			DB:       0,
		},
	}
	testIncome = NewIncome(testIncomeConf, onchain.NewService(), chainSvc.NewService(testChainConf), wallet.NewService(testWalletConf))
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
			txn.GetId(),
			txn.GetChain(),
			txn.GetAsset(),
			txn.GetTxHash(),
			txn.GetSender(),
			txn.GetReceiver(),
			txn.GetAmountRaw(),
			txn.GetPrecession(),
			txn.GetAmount(),
			txn.GetVOut(),
			txn.GetStatus()))
	}
}
