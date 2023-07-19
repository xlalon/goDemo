package conf

import (
	"flag"

	onchainConf "github.com/xlalon/golee/internal/xchain/conf"
	"github.com/xlalon/golee/pkg/database/mysql"
	"github.com/xlalon/golee/pkg/database/redis"
	"github.com/xlalon/golee/pkg/json"
	"github.com/xlalon/golee/pkg/net/http/server"
)

var (
	Conf = &Config{
		Server: &server.Config{
			Debug:   true,
			Address: ":8080",
		},
	}

	confPath string
)

func init() {
	flag.StringVar(&confPath, "conf", "config.yaml", "config path")
}

func Init() error {

	var err error

	Conf, err = fromFile(confPath)

	json.JPrint("Config", Conf)

	return err
}

type Config struct {
	Server *server.Config `yaml:"server"`

	Mysql *mysql.Config `yaml:"mysql"`
	Redis *redis.Config `yaml:"redis"`

	Chain *onchainConf.Config `yaml:"chain"`
}
