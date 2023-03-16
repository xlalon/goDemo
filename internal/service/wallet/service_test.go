package wallet

import (
	onchainConf "github.com/xlalon/golee/internal/onchain/conf"
	"github.com/xlalon/golee/internal/onchain/x"
	"github.com/xlalon/golee/internal/service/wallet/conf"
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

	testConf = &conf.Config{
		Mysql: &mysql.Config{DNS: "mycat:p123456@tcp(127.0.0.1:3306)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"},
		Redis: &redis.Config{
			Address:  "127.0.0.1",
			Port:     6379,
			Password: "",
			DB:       0,
		},
	}
	testSvc = NewService(testConf)
)

//func TestService_NewAccount(t *testing.T) {
//	acct, err := testSvc.NewAccount("BAND", "DEPOSIT")
//	if err != nil {
//		t.Error(err)
//	}
//	json.PPrint("new account", acct)
//}

//func TestService_GetAccount(t *testing.T) {
//	acct := domain.AccountFactory(&domain.AccountDTO{
//		Id:      6627670032386,
//		Chain:   "BAND",
//		Address: "band1ggq8us6lh4c8hr4624xnrlud6q5lqhklakysnd",
//		Memo:    "614322816608",
//		Label:   "HOT",
//		Status:  "VALID",
//	})
//	_ = testSvc.walletRepo.Save(acct)
//}
