package api

import (
	"encoding/json"
	"net/http"
	"olympic-medals-table/models"
)

func GetCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	models.Position()

	for i := 0; i < len(models.AllCountries); i++ {
		models.AllCountries[i].AllMedals = models.AllCountries[i].GoldMedals +
			models.AllCountries[i].SilverMedals + models.AllCountries[i].BronzeMedals
		models.AllCountries[i].Position = i + 1
	}
	json.NewEncoder(w).Encode(models.AllCountries)
}
