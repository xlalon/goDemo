package conf

import (
	"flag"

	"github.com/xlalon/golee/internal/infra/repository"
	"github.com/xlalon/golee/internal/infra/repository/chain"
	"github.com/xlalon/golee/internal/infra/repository/deposit"
	"github.com/xlalon/golee/internal/infra/repository/wallet"
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
	Conf.Repository = &repository.Config{
		Chain: &chain.Config{
			Mysql: Conf.Mysql,
			Redis: Conf.Redis,
		},
		Deposit: &deposit.Config{
			Mysql: Conf.Mysql,
			Redis: Conf.Redis,
		},
		Wallet: &wallet.Config{
			Mysql: Conf.Mysql,
			Redis: Conf.Redis,
		},
	}

	json.PPrint("Config", Conf)

	return err
}

type Config struct {
	Repository *repository.Config `yaml:"repository"`
	Mysql      *mysql.Config      `yaml:"mysql"`
	Redis      *redis.Config      `yaml:"redis"`

	Chain *conf.Config `yaml:"chain"`

	Job *worker.Config `yaml:"job"`
}
