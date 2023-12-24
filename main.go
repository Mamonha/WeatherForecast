package main

import (
	"net/http"
	"weather/controllers"
)

func main() {
	http.HandleFunc("/location", controllers.GetWeatherData)
	http.ListenAndServe(":8080", nil)
}
