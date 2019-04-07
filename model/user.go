package model

type User struct {
	Id       uint `gorm:"primary_key"`
	Nama     string
	Password string
}
