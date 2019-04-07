package db

import (
	"rest_api/model"
)

func Migrate() {
	db := InitDb()
	db.SingularTable(true)
	db.AutoMigrate(&model.Mahasiswa{})
}
