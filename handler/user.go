package handler

import (
	"encoding/json"
	"net/http"

	"go_jwt_redis/db"
	"go_jwt_redis/model"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user model.Users
	json.NewDecoder(r.Body).Decode(&user)
	db.DBCon.Create(&user)
	json.NewEncoder(w).Encode(user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user model.Users
	id := r.URL.Query().Get("id")

	db.DBCon.First(&user, id)
	db.DBCon.Delete(&user)
	json.NewEncoder(w).Encode(user)

}
