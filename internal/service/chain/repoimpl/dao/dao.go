package dao

import (
	"github.com/xlalon/golee/internal/service/chain/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dao struct {
	db *gorm.DB
}

func New(conf *conf.Config) (d *Dao) {

	mysqlDb, err := gorm.Open(mysql.Open(conf.Mysql.DNS), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}
	return &Dao{
		db: mysqlDb,
	}
}
