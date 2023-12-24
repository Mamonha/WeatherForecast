package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"weather/models/api"

	"github.com/joho/godotenv"
)

func GetWeatherData(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	apiKey := os.Getenv("API_KEY")
	endpoint := os.Getenv("URL_LOCATION")
	lon := "-25.503212"
	lat := "-54.572352"
	url := fmt.Sprintf(endpoint, lat, lon, apiKey)

	response, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data api.WeatherData
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}
