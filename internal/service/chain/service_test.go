package chain

import (
	"fmt"
	"testing"

	"github.com/xlalon/golee/internal/service/chain/conf"
	"github.com/xlalon/golee/pkg/database/mysql"
	"github.com/xlalon/golee/pkg/database/redis"
	"github.com/xlalon/golee/pkg/json"
)

var (
	testConf = &conf.Config{
		Mysql: &mysql.Config{DNS: "mycat:p123456@tcp(127.0.0.1:8066)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"},
		Redis: &redis.Config{
			Address:  "127.0.0.1",
			Port:     6379,
			Password: "",
			DB:       0,
		},
	}
)

//func TestChain_Save(t *testing.T) {
//	c := &model.Chain{
//		Code: "WAX",
//		Name: "Worldwide Asset eXchange",
//	}
//	chain := InitChain(c, testConf)
//	err := chain.Save()
//	if err != nil {
//		fmt.Println("err", err)
//	}
//}

func TestChain_GetAssets(t *testing.T) {
	svc := NewService(testConf)
	assets, err := svc.GetAssets()
	if err != nil {
		fmt.Println("err", err)
	}
	json.PPrint("assets", assets)
}

//func TestChain_GetLatestHeight(t *testing.T) {
//	c := &model.Chain{
//		Code: "BAND",
//		Name: "Band",
//	}
//	chain := InitChain(c, testConf)
//	fmt.Println("height", chain.GetLatestHeight())
//}
