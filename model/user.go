package model

type Users struct {
	Id       uint `gorm:"primary_key"`
	Nama     string
	Password string
}
