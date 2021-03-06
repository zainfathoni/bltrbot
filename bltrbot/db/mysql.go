package db

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func MysqlDB() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("MYSQL_DATABASE_USER")+":"+os.Getenv("MYSQL_DATABASE_PASSWORD")+"@tcp("+os.Getenv("MYSQL_DATABASE_HOST")+":"+os.Getenv("MYSQL_DATABASE_PORT")+")/"+os.Getenv("MYSQL_DATABASE_NAME")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}
