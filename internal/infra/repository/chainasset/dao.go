package chainasset

import (
	xmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/xlalon/golee/pkg/database/mysql"
	"github.com/xlalon/golee/pkg/database/redis"
)

type Config struct {
	Mysql *mysql.Config
	Redis *redis.Config
}

type Dao struct {
	db *gorm.DB
}

func NewDao(conf *Config) *Dao {

	mysqlDb, err := gorm.Open(xmysql.Open(conf.Mysql.DNS), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}
	return &Dao{
		db: mysqlDb,
	}
}

func (d *Dao) NextId() int64 {
	return mysql.NextID()
}
