package main

import (
	"context"
	"fmt"
	"log"
	"async-weather/openmeteo"
)

func main() {
	// This is temp test implementation. TODO: http server with stream of 
	lat := 50.05
	lng := 21.41

	ctx := context.Background()

	weatherData, err := openmeteo.GetCurrentWeather(ctx, lat, lng)
	if err != nil {
		log.Fatalf("Error fetching weather data: %v", err)
	}

	fmt.Printf("Temperature: %.1f %s\n", weatherData.Current.Temperature2M, weatherData.CurrentUnits.Temperature2M)
}
