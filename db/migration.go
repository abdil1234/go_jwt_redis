package db

import (
	"github.com/abdil1234/go_jwt_redis/model"
)

func Migrate() {
	DBCon.SingularTable(true)
	DBCon.AutoMigrate(&model.Mahasiswa{}, model.Users{})
}
