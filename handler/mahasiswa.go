package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest_api/db"
	"rest_api/model"
	"time"

	"github.com/go-redis/redis"
)

func AllMahasiswa(w http.ResponseWriter, r *http.Request) {
	mysql := db.InitDb()
	defer mysql.Close()

	redisClient := db.InitRedis()
	defer redisClient.Close()

	var mahasiswas []model.Mahasiswa

	//get redis data
	mahasiswaData, err := redisClient.Get("mahasiswas").Result()

	if err == redis.Nil {
		mysql.Find(&mahasiswas)
		data, _ := json.Marshal(mahasiswas)

		//set redis data
		err := redisClient.Set("mahasiswas", data, 3*time.Second).Err()
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	fmt.Println(w, mahasiswaData)

}

func CreateMahasiswa(w http.ResponseWriter, r *http.Request) {
	db := db.InitDb()
	defer db.Close()

	var mahasiswa model.Mahasiswa
	json.NewDecoder(r.Body).Decode(&mahasiswa)
	db.Create(&mahasiswa)
	json.NewEncoder(w).Encode(mahasiswa)

}

func UpdateMahasiswa(w http.ResponseWriter, r *http.Request) {
	mysql := db.InitDb()
	defer mysql.Close()

	redisClient := db.InitRedis()
	defer redisClient.Close()

	var mahasiswa model.Mahasiswa
	id := r.URL.Query().Get("id")
	key := "mahasiswa:" + string(id)

	//get redis data
	mahasiswaData, err := redisClient.Get(key).Result()

	if err == redis.Nil {
		mysql.First(&mahasiswa, id)
		data, _ := json.Marshal(mahasiswa)

		//set redis data
		err := redisClient.Set(key, data, 5*time.Second).Err()
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	json.Unmarshal([]byte(mahasiswaData), &mahasiswa)

	json.NewDecoder(r.Body).Decode(&mahasiswa)

	mysql.Save(&mahasiswa)
	json.NewEncoder(w).Encode(mahasiswa)
}

func DeleteMahasiswa(w http.ResponseWriter, r *http.Request) {

	mysql := db.InitDb()
	defer mysql.Close()

	redisClient := db.InitRedis()
	defer redisClient.Close()

	var mahasiswa model.Mahasiswa
	id := r.URL.Query().Get("id")
	key := "mahasiswa:" + string(id)

	mahasiswaData, err := redisClient.Get(key).Result()

	if err == redis.Nil {
		mysql.First(&mahasiswa, id)
		data, _ := json.Marshal(mahasiswa)

		//set redis data
		err := redisClient.Set(key, data, 5*time.Second).Err()
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	json.Unmarshal([]byte(mahasiswaData), &mahasiswa)

	mysql.Delete(&mahasiswa)

	json.NewEncoder(w).Encode(mahasiswa)

}
