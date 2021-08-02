package api

import (
	"encoding/json"
	"net/http"
)

func GetCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	position()

	for i := 0; i < len(allCountries); i++ {
		allCountries[i].AllMedals = allCountries[i].GoldMedals +
			allCountries[i].SilverMedals + allCountries[i].BronzeMedals
		allCountries[i].Position = i + 1
	}
	json.NewEncoder(w).Encode(allCountries)
}
