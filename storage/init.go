package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vinoMamba/lazy-doc-end/config"
)

var db *sql.DB

func init() {
}

func DbConn() {
	var err error
	mysqlConfig := config.GetMysqlConfig()
	dns := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",
		mysqlConfig.Username,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.Database,
	)
	db, err = sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func NewQuery() *Queries {
	return New(db)
}
