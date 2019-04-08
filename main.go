package main

import (
	"fmt"
	"log"
	"net/http"
	"rest_api/db"
	"rest_api/handler"
	"rest_api/middleware"

	"github.com/gorilla/mux"
)

var port string = ":8081"

func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)

	//gorilla middleware
	r.Use(middleware.MiddlewareJWTAuthorization)

	//auth route
	r.HandleFunc("/login", handler.HandlerLogin)

	//mahasiswa route
	r.HandleFunc("/mahasiswas", handler.AllMahasiswa).Methods("GET")
	r.HandleFunc("/mahasiswas/create", handler.CreateMahasiswa).Methods("POST")
	r.HandleFunc("/mahasiswas/update/{id}", handler.UpdateMahasiswa).Methods("PATCH")
	r.HandleFunc("/mahasiswas/delete/{id}", handler.DeleteMahasiswa).Methods("DELETE")

	//user route
	r.HandleFunc("/users/create", handler.CreateUser).Methods("POST")
	r.HandleFunc("/users/delete/{id}", handler.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(port, r))
}

func main() {
	fmt.Println("Go run on port %s \n", port)
	db.Migrate()
	handleRequest()
}
