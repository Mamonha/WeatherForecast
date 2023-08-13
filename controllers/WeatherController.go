package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"weather/models/api"
)

func GetWeatherData() api.WeatherData {
	apiKey := "915bc9171eddaec8223211d99469e0b0"
	lon := "-25.503212"
	lat := "-54.572352"

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s", lat, lon, apiKey)

	response, err := http.Get(url)
	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var data api.WeatherData

	err2 := json.Unmarshal(body, &data)

	if err2 != nil {
		log.Fatal(err2)
	}

	return data

}

func ValidateData(data *api.WeatherData) {

}
