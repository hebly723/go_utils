package utils

import (
	"fmt"
	"log"

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

var DB *gorm.DB

func InitDB(dbc Db) error {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		dbc.User, dbc.Password, dbc.Host, dbc.Port, dbc.Database)
	log.Println(dsn)
	var err error
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)}); err != nil {
		return err
	}

	return nil
}
