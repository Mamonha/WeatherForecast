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

func GetWeatherData() api.WeatherData {
	err := godotenv.Load(".env")
	apiKey := os.Getenv("API_KEY")

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
	fmt.Println(data)
	return data

}

// func ValidateData(data *api.WeatherData) {

// }
