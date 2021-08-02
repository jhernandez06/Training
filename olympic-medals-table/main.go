package main

import (
	"log"
	"net/http"
	"olympic-medals-table/actions"
	"olympic-medals-table/api"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/api/countries", api.GetCountries).Methods("GET")
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	r.HandleFunc("/", actions.FileServer)

	server := &http.Server{
		Addr:           ":3000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening...")
	server.ListenAndServe()

}
