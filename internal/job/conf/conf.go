package conf

import (
	"flag"

	"github.com/xlalon/golee/internal/onchain/conf"
	"github.com/xlalon/golee/pkg/database/mysql"
	"github.com/xlalon/golee/pkg/database/redis"
	"github.com/xlalon/golee/pkg/job/worker"
	"github.com/xlalon/golee/pkg/json"
)

var (
	Conf = &Config{}

	confPath string
)

func init() {
	flag.StringVar(&confPath, "conf", "config.yaml", "config path")
}

func Init() error {
	var err error
	Conf, err = fromFile(confPath)

	json.PPrint("Config", Conf)

	return err
}

type Config struct {
	Mysql *mysql.Config `yaml:"mysql"`
	Redis *redis.Config `yaml:"redis"`

	Chain *conf.Config `yaml:"chain"`

	Job *worker.Config `yaml:"job"`
}
