package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"olympic-medals-table/models"
	"time"

	"github.com/gorilla/mux"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}
func getCountrys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(models.AllCountrys)
}

func main() {

	r := mux.NewRouter().StrictSlash(true)
	//r.HandleFunc("/", indexRoute)
	r.HandleFunc("/countrys", getCountrys).Methods("GET")

	r.PathPrefix("/web-server/public/").Handler(http.StripPrefix("/web-server/public/", http.FileServer(http.Dir("./web-server/public"))))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := ioutil.ReadFile("web-server/public/index.html")
		if err != nil {
			panic(err)
		}
		w.Write(file)
	})
	fmt.Println("Runing...")

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
