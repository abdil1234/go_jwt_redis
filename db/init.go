package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:kerjakansekarang@/rest_golang?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed connect db")
	}

	return db
}
