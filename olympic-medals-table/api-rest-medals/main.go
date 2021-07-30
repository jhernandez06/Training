package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"api-rest-medals/models"

	"github.com/gorilla/mux"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}
func getCountrys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(models.AllCountrys)
}

func getCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryselec := vars["country"]

	for _, country := range models.AllCountrys {
		if strings.ToLower(country.Country) == strings.ToLower(countryselec) {
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(country)
		} else {
			fmt.Fprintf(w, "invalid Country")
			return
		}
	}
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/countrys", getCountrys).Methods("GET")
	router.HandleFunc("/countrys/{country}", getCountry).Methods("GET")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening...")
	server.ListenAndServe()

}
