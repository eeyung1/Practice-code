package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		Windspeed   float64 `json:"windspeed"`
		Weathercode int     `json:"weathercode"`
	} `json:"current_weather"`
}

func main() {
	url := "https://api.open-meteo.com/v1/forecast?latitude=6.5244&longitude=3.3792&current_weather=true"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// parse JSON into struct
	var weather WeatherResponse
	json.Unmarshal(body, &weather)

	// print the results
	fmt.Printf("Temperature: %.1f°C\n", weather.CurrentWeather.Temperature)
	fmt.Printf("Windspeed:   %.1f km/h\n", weather.CurrentWeather.Windspeed)
}