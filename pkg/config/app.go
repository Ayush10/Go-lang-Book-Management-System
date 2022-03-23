package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(mysql.Open("Ayush:ayushlee10@tcp(127.0.0.1:3306/simplebookstore?charset=utf8mb4&parseTime=true&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
