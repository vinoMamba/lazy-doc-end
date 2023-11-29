package storage

import (
	"fmt"

	"github.com/vinoMamba/lazy-doc-end/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

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
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = DB.Exec("select 1;").Error
	if err != nil {
		panic(err)
	}
}
