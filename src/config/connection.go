package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	host := "127.0.0.1"
	port := "3306"
	dbName := "skyshi"
	username := "root"
	password := "M@utauaja982"

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8&parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Tidak dapat terkoneksi dengan database, periksa config!")
	}

	return db
}
