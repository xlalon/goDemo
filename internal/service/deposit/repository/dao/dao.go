package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/xlalon/golee/internal/service/deposit/conf"
)

type Dao struct {
	db *gorm.DB
}

func New(conf *conf.Config) (d *Dao) {

	mysqlDb, err := gorm.Open(mysql.Open(conf.Mysql.DNS), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Dao{
		db: mysqlDb,
	}
}
