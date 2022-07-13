package database

import (
	"crawl/util"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConn(config util.Config) (db *gorm.DB) {
	dns := fmt.Sprintf(
		"%s:%s@(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBUserName,
		config.DBPassWord,
		config.DBHost,
		config.DBName)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic("failed to connection db")
	}
	return db
}
