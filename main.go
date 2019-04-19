package main

import (
	"fmt"
	"log"
	"net/http"
	"rest_api/db"
	"rest_api/handler"
	"rest_api/middleware"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

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

	log.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	var err error

	//db connection
	db.DBCon, err = gorm.Open("mysql", "root:kerjakansekarang@/rest_golang?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed connect db")
	}

	defer db.DBCon.Close()

	//redis connection
	db.ReCon = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:4321",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = db.ReCon.Ping().Result()
	defer db.ReCon.Close()

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed connect redis")
	}

	fmt.Println("Go run on port %s \n", ":8081")
	db.Migrate()
	handleRequest()
}
