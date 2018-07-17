package main

import (
	"github.com/gorilla/mux"
	. "rest-api/config"
	. "rest-api/handler"
	. "rest-api/db"
	"log"
	"net/http"
)

var config = Config{}
var dao = CompaniesDAO{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api/v1").Subrouter()

	sub.HandleFunc("/companies", AllCompaniesEndPoint).Methods("GET")
	sub.HandleFunc("/companies", CreateMovieEndPoint).Methods("POST")
	sub.HandleFunc("/companies", UpdateCompanyEndpoint).Methods("PUT")
	sub.HandleFunc("/companies", DeleteCompanyEndpoint).Methods("DELETE")
	sub.HandleFunc("/companies/{name}", FindCompanyEndpoint).Methods("GET")

	if err := http.ListenAndServe(":3000", sub); err != nil {
		log.Fatal(err)
	}
}