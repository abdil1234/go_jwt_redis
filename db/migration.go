package db

import (
	"go_jwt_redis/model"
)

func Migrate() {
	DBCon.SingularTable(true)
	DBCon.AutoMigrate(&model.Mahasiswa{}, model.Users{})
}
