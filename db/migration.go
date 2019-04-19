package db

import (
	"rest_api/model"
)

func Migrate() {
	DBCon.SingularTable(true)
	DBCon.AutoMigrate(&model.Mahasiswa{}, model.Users{})
}
