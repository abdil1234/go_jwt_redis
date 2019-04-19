package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"rest_api/db"
	"rest_api/model"

	"github.com/go-redis/redis"
)

func AllMahasiswa(w http.ResponseWriter, r *http.Request) {

	var mahasiswas []model.Mahasiswa

	//get redis data
	mahasiswaData, err := db.ReCon.Get("mahasiswas").Result()

	if err == redis.Nil {
		db.DBCon.Find(&mahasiswas)
		data, _ := json.Marshal(mahasiswas)

		//set redis data
		err := db.ReCon.Set("mahasiswas", data, 3*time.Second).Err()
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	fmt.Fprint(w, mahasiswaData)

}

func CreateMahasiswa(w http.ResponseWriter, r *http.Request) {

	var mahasiswa model.Mahasiswa
	json.NewDecoder(r.Body).Decode(&mahasiswa)
	db.DBCon.Create(&mahasiswa)
	json.NewEncoder(w).Encode(mahasiswa)

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
	json.NewEncoder(w).Encode(mahasiswa)
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

	json.NewEncoder(w).Encode(mahasiswa)

}
