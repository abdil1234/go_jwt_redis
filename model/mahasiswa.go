package model

import (
	"time"
)

type Mahasiswa struct {
	Id      uint `gorm:"primary_key"`
	Nama    string
	Nip     string
	Tgl_lhr time.Time
	Alamat  string
	Email   string
}
