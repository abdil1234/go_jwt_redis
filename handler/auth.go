package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"go_jwt_redis/db"
	"go_jwt_redis/helpers"
	"go_jwt_redis/model"

	jwt "github.com/dgrijalva/jwt-go"
)

func HandlerLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Unsupported http method", http.StatusBadRequest)
		return
	}

	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	ok, nama := authenticateUser(username, password)
	if !ok {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	claims := helpers.MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    helpers.APPLICATION_NAME,
			ExpiresAt: time.Now().Add(helpers.LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Username: nama,
	}

	token := jwt.NewWithClaims(
		helpers.JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(helpers.JWT_SIGNATURE_KEY)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenString, _ := json.Marshal(helpers.M{"token": signedToken})
	w.Write([]byte(tokenString))
}

func authenticateUser(username, password string) (bool, string) {
	var user model.Users
	var emptyUser = model.Users{}

	db.DBCon.Where("nama = ? AND password = ?", username, password).Find(&user)

	if user != emptyUser {
		return true, username
	}

	return false, ""
}
