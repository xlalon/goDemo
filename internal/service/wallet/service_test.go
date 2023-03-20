package wallet

import (
	"github.com/xlalon/golee/internal/infra/repository"
	"github.com/xlalon/golee/internal/infra/repository/wallet"
	onchainConf "github.com/xlalon/golee/internal/onchain/conf"
	"github.com/xlalon/golee/internal/onchain/x"
	"github.com/xlalon/golee/internal/service/wallet/domain"
	"github.com/xlalon/golee/pkg/database/mysql"
	"github.com/xlalon/golee/pkg/database/redis"
	"testing"
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
	_             = x.Init(testOnchainConf)
	testMysqlConf = &mysql.Config{DNS: "root:Xiao0000@tcp(127.0.0.1:3306)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"}
	testRedisConf = &redis.Config{
		Address:  "127.0.0.1",
		Port:     6379,
		Password: "",
		DB:       0,
	}

	testConf = &repository.Config{
		Wallet: &wallet.Config{
			Mysql: testMysqlConf,
			Redis: testRedisConf,
		},
	}
	testRepository = repository.NewRegistry(testConf)
	testSvc        = NewService(testRepository.WalletRepository())
)

func TestService_NewAccount(t *testing.T) {
	acct := domain.AccountFactory(&domain.AccountDTO{
		Id:      mysql.NextID(),
		Chain:   "BAND",
		Address: "band1dkl8wga94803qygwdspwa5kxdfyjpt8zr0uzh9",
		Memo:    "714322816608",
		Label:   "DEPOSIT",
		Status:  "VALID",
	})
	_ = testSvc.walletRepo.Save(acct)
}
