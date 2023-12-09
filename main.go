package main

import (
	"fmt"
	"weather/controllers"
)

func main() {

	fmt.Print("Teste")
	controllers.GetWeatherData()
	controllers.ValidateData()
}
