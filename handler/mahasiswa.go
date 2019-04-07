package handler

import (
	"encoding/json"
	"net/http"
	"rest_api/db"
	"rest_api/model"
)

func AllMahasiswa(w http.ResponseWriter, r *http.Request) {
	db := db.InitDb()
	defer db.Close()

	var mahasiswas []model.Mahasiswa
	db.Find(&mahasiswas)

	json.NewEncoder(w).Encode(mahasiswas)
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
	db := db.InitDb()
	defer db.Close()
	var mahasiswa model.Mahasiswa
	id := r.URL.Query().Get("id")
	db.First(&mahasiswa, id)
	json.NewDecoder(r.Body).Decode(&mahasiswa)

	db.Save(&mahasiswa)
	json.NewEncoder(w).Encode(mahasiswa)
}

func DeleteMahasiswa(w http.ResponseWriter, r *http.Request) {
	db := db.InitDb()
	defer db.Close()
	var mahasiswa model.Mahasiswa
	id := r.URL.Query().Get("id")

	db.First(&mahasiswa, id)
	db.Delete(&mahasiswa)
	json.NewEncoder(w).Encode(mahasiswa)

}
