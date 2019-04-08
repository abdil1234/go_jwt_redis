package handler

import (
	"encoding/json"
	"net/http"
	"rest_api/db"
	"rest_api/model"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := db.InitDb()
	defer db.Close()

	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	json.NewEncoder(w).Encode(user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := db.InitDb()
	defer db.Close()
	var user model.User
	id := r.URL.Query().Get("id")

	db.First(&user, id)
	db.Delete(&user)
	json.NewEncoder(w).Encode(user)

}
