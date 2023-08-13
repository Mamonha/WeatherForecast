package main

import (
	"weather/controllers"
)

func main() {
	controllers.GetWeatherData()
	controllers.ValidateData()
}
