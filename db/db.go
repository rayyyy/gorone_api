package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	dsn := "root:pass@tcp(db:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB Connection Error")
	}
}

func DbManager() *gorm.DB {
	return db
}
