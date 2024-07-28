// api/regencies.go

package handler

import (
	"encoding/json"
	"net/http"
)

type Regency struct {
	Code         string `json:"code"`
	ProvinceCode string `json:"province_code"`
	Name         string `json:"name"`
}

func RegenciesHandler(w http.ResponseWriter, r *http.Request) {
	provinceCode := r.URL.Query().Get("province_code")
	if provinceCode == "" {
		http.Error(w, "Province code is required", http.StatusBadRequest)
		return
	}

	// Contoh data statis, biasanya Anda akan mengambil data dari sumber yang sesuai
	regencies := []Regency{
		{Code: "36.01", ProvinceCode: "36", Name: "KABUPATEN PANDEGLANG"},
		{Code: "36.02", ProvinceCode: "36", Name: "KABUPATEN LEBAK"},
	}

	// Filter berdasarkan province_code
	var filteredRegencies []Regency
	for _, regency := range regencies {
		if regency.ProvinceCode == provinceCode {
			filteredRegencies = append(filteredRegencies, regency)
		}
	}

	response := struct {
		Data []Regency `json:"data"`
	}{
		Data: filteredRegencies,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
