package conf

import (
	"github.com/xlalon/golee/pkg/database/mysql"
	"github.com/xlalon/golee/pkg/database/redis"
)

type Config struct {
	Mysql *mysql.Config
	Redis *redis.Config
}
