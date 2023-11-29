package storage

import (
	"fmt"

	"github.com/vinoMamba/lazy-doc-end/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func DbConn() {
	var err error
	mysqlConfig := config.GetMysqlConfig()
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.Username,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.Database,
	)
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.Exec("select 1;").Error
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
