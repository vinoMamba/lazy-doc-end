package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySqlOpts struct {
	Host     string
	Username string
	Password string
	Database string
	LogLevel int
}

func (opts *MySqlOpts) DSN() string {
	return fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		opts.Username,
		opts.Password,
		opts.Host,
		opts.Database,
	)
}

func NewMySql(opts *MySqlOpts) (*gorm.DB, error) {

	logLevel := logger.Silent
	if opts.LogLevel != 0 {
		logLevel = logger.LogLevel(opts.LogLevel)
	}

	db, err := gorm.Open(mysql.Open(opts.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
