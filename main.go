package main

import (
	"fmt"
	"log"
	"net/http"
	"rest_api/db"
	"rest_api/handler"

	"github.com/gorilla/mux"
)

var port string = ":8081"

func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/mahasiswas", handler.AllMahasiswa).Methods("GET")
	r.HandleFunc("/mahasiswas/create", handler.CreateMahasiswa).Methods("POST")
	r.HandleFunc("/mahasiswas/update/{id}", handler.UpdateMahasiswa).Methods("PATCH")
	r.HandleFunc("/mahasiswas/delete/{id}", handler.DeleteMahasiswa).Methods("DELETE")
	log.Fatal(http.ListenAndServe(port, r))
}

func main() {
	fmt.Println("Go run on port %s \n", port)
	db.Migrate()
	handleRequest()
}
