package main

import (
	"fmt"
)

type Fetcher interface {
	Fetch() string
}

type OpenWeatherMap struct{}

func (ow OpenWeatherMap) Fetch() string {
	return "Weather data from OpenWeatherMap: 22°C"
}

type WeatherAPI struct{}

func (wa WeatherAPI) Fetch() string {
	return "Weather data from WeatherAPI: Sunny, 24°C"
}

func ProcessWeather(f Fetcher) {
	data := f.Fetch()
	fmt.Println(data)
}

func main(){
	ow := OpenWeatherMap{}
	wa := WeatherAPI{}


	ProcessWeather(ow)
	ProcessWeather(wa)
}