package chain

import (
	"github.com/xlalon/golee/internal/infra/repository"
	"github.com/xlalon/golee/internal/infra/repository/chain"
	"github.com/xlalon/golee/pkg/database/mysql"
	"github.com/xlalon/golee/pkg/database/redis"
)

var (
	testMysqlConf = &mysql.Config{DNS: "root:Xiao0000@tcp(127.0.0.1:3306)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"}
	testRedisConf = &redis.Config{
		Address:  "127.0.0.1",
		Port:     6379,
		Password: "",
		DB:       0,
	}

	testRepository = repository.NewRegistry(&repository.Config{
		Chain: &chain.Config{
			Mysql: testMysqlConf,
			Redis: testRedisConf,
		},
	})
	testSvc = NewService(testRepository.ChainRepository())
)

//func TestService_NewChain(t *testing.T) {
//	chain := &domain.ChainDTO{
//		Id:     6627923918848,
//		Code:   "BTC",
//		Name:   "BTC",
//		Status: string(domain.ChainStatusOnline),
//		Assets: []*domain.AssetDTO{{
//			Id:         mysql.NextID(),
//			Code:       "BTC2",
//			Name:       "BTC2",
//			Chain:      "BTC",
//			Identity:   "btc2",
//			Precession: 8,
//			Status:     string(domain.ChainStatusOnline),
//		},
//			{
//				Id:         6627986571266,
//				Code:       "BTC3",
//				Name:       "BTC3",
//				Chain:      "BTC",
//				Identity:   "btc3",
//				Precession: 8,
//				Status:     string(domain.ChainStatusOnline),
//			}},
//	}
//	json.PPrint("chain", chain)
//	testSvc.NewChain(chain)
//}

//func TestService_GetChains(t *testing.T) {
//	chains, err := testSvc.GetChains()
//	if err != nil {
//		fmt.Println("GetChains err", err)
//	}
//	json.PPrint("chains", chains)
//
//}
//
//func TestService_GetAssets(t *testing.T) {
//	assets, err := testSvc.GetAssets()
//	if err != nil {
//		fmt.Println("GetAssets err", err)
//	}
//	json.PPrint("assets", assets)
//}
