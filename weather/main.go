package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"net/url"
)

type WeatherResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		Windspeed   float64 `json:"windspeed"`
		Weathercode int     `json:"weathercode"`
	} `json:"current_weather"`
}

type GeoResponse struct {
	Results []struct {
		Name      string  `json:"name"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Country   string  `json:"country"`
	} `json:"results"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . Lagos")
		os.Exit(1)
	}

	city := os.Args[1]

	// FIRST API CALL — geocoding
	geoURL := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1", url.QueryEscape(city))
	resp, err := http.Get(geoURL)
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

	// parse geo response
	var geo GeoResponse
	json.Unmarshal(body, &geo)

	// check city exists
	if len(geo.Results) == 0 {
		fmt.Println("City not found!")
		os.Exit(1)
	}

	// extract coordinates
	lat := geo.Results[0].Latitude
	lon := geo.Results[0].Longitude
	name := geo.Results[0].Name
	country := geo.Results[0].Country

	// SECOND API CALL — weather
	weatherURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true", lat, lon)
	resp2, err := http.Get(weatherURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp2.Body.Close()

	body2, err := io.ReadAll(resp2.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// parse weather response
	var weather WeatherResponse
	json.Unmarshal(body2, &weather)

	// print results
	fmt.Printf("City:        %s, %s\n", name, country)
	fmt.Printf("Temperature: %.1f°C\n", weather.CurrentWeather.Temperature)
	fmt.Printf("Windspeed:   %.1f km/h\n", weather.CurrentWeather.Windspeed)
}
