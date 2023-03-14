package deposit

import (
	"fmt"
	"testing"

	"github.com/xlalon/golee/internal/service/deposit/conf"
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
	testSvc = NewService(testConf)
)

func TestService_GetDeposit(t *testing.T) {
	deps, err := testSvc.GetDeposits()
	if err != nil {
		fmt.Println("GetDeposit err:", err)
	}
	json.PPrint("deposit", deps)
}
