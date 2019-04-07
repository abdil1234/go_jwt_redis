package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"rest_api/db"
	"rest_api/handler"
	"rest_api/helpers"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var port string = ":8081"

func MiddlewareJWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != helpers.JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return helpers.JWT_SIGNATURE_KEY, nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(context.Background(), "username", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})

}
func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)

	//gorilla middleware
	r.Use(MiddlewareJWTAuthorization)

	//auth route
	r.HandleFunc("/login", handler.HandlerLogin)

	//mahasiswa route
	r.HandleFunc("/mahasiswas", handler.AllMahasiswa).Methods("GET")
	r.HandleFunc("/mahasiswas/create", handler.CreateMahasiswa).Methods("POST")
	r.HandleFunc("/mahasiswas/update/{id}", handler.UpdateMahasiswa).Methods("PATCH")
	r.HandleFunc("/mahasiswas/delete/{id}", handler.DeleteMahasiswa).Methods("DELETE")

	//user route
	r.HandleFunc("/users/create", handler.CreateUser).Methods("POST")
	r.HandleFunc("/users/update/{id}", handler.UpdateUser).Methods("PATCH")
	r.HandleFunc("/users/delete/{id}", handler.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(port, r))
}

func main() {
	fmt.Println("Go run on port %s \n", port)
	db.Migrate()
	handleRequest()
}
