package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/abdil1234/go_jwt_redis/db"
	"github.com/abdil1234/go_jwt_redis/helpers"
	"github.com/abdil1234/go_jwt_redis/model"

	"github.com/go-redis/redis"
)

func AllMahasiswa(w http.ResponseWriter, r *http.Request) {

	var mahasiswas []model.Mahasiswa
	//cek redis data
	_, err := db.ReCon.Get("mahasiswas").Result()

	if err == redis.Nil {

		db.DBCon.Find(&mahasiswas)
		data, _ := json.Marshal(mahasiswas)
		//set redis data
		err := db.ReCon.Set("mahasiswas", data, time.Hour).Err()
		if err != nil {
			panic(err)
		}
	}

	//get redis data

	mahasiswaData, _ := db.ReCon.Get("mahasiswas").Result()

	helpers.ResJSON(w, []byte(mahasiswaData))

}

func CreateMahasiswa(w http.ResponseWriter, r *http.Request) {

	var mahasiswa model.Mahasiswa
	json.NewDecoder(r.Body).Decode(&mahasiswa)
	db.DBCon.Create(&mahasiswa)

	d, _ := json.Marshal(mahasiswa)

	//empty redis data
	db.ReCon.Del("mahasiswas")

	helpers.ResJSON(w, []byte(d))

}

func UpdateMahasiswa(w http.ResponseWriter, r *http.Request) {

	var mahasiswa model.Mahasiswa
	id := r.URL.Query().Get("id")
	key := "mahasiswa:" + string(id)

	//get redis data
	mahasiswaData, err := db.ReCon.Get(key).Result()

	if err == redis.Nil {
		db.DBCon.First(&mahasiswa, id)
		data, _ := json.Marshal(mahasiswa)

		//set redis data
		err := db.ReCon.Set(key, data, 5*time.Second).Err()
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	json.Unmarshal([]byte(mahasiswaData), &mahasiswa)

	json.NewDecoder(r.Body).Decode(&mahasiswa)

	db.DBCon.Save(&mahasiswa)
	d, _ := json.Marshal(mahasiswa)

	helpers.ResJSON(w, []byte(d))
}

func DeleteMahasiswa(w http.ResponseWriter, r *http.Request) {

	var mahasiswa model.Mahasiswa
	id := r.URL.Query().Get("id")
	key := "mahasiswa:" + string(id)

	mahasiswaData, err := db.ReCon.Get(key).Result()

	if err == redis.Nil {
		db.DBCon.First(&mahasiswa, id)
		data, _ := json.Marshal(mahasiswa)

		//set redis data
		err := db.ReCon.Set(key, data, 5*time.Second).Err()
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	json.Unmarshal([]byte(mahasiswaData), &mahasiswa)
	db.DBCon.Delete(&mahasiswa)

	d, _ := json.Marshal(mahasiswa)

	helpers.ResJSON(w, []byte(d))

}
