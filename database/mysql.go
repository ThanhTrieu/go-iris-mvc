package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectSQL(host, port, uname, pass, dbname string) (*gorm.DB, error) {

	dsn := uname + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
