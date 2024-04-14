package utils

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Db struct {
	Database string
	Host     string
	Port     string
	Password string
	User     string
}

func InitDB(dbc Db) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(
		fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			dbc.User, dbc.Password, dbc.Host, dbc.Port, dbc.Database)),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
}
