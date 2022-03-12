package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func ConnectGORM() *gorm.DB {
	DBMS := "mysql"
	USER := "user"
	PASS := "password"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "sample_db"

	CONECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
