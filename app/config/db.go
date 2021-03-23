package config

import (
	"github.com/jinzhu/gorm"
	"github.com/sho0126hiro/toriton/app/internal/interface/dao"
)

func DBConnect() (*gorm.DB, error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "####"
	PROTOCOL := "tcp(##.###.##.###:3306)"
	DBNAME := "##"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	return gorm.Open(DBMS, CONNECT)
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&dao.User{})
	return db
}
