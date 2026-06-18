package main

import "fmt"

type Artist struct {
	ID int
	Name string
}

type Location struct {
	ID int
	Locations []string
}

type ArtistInfo struct {
	Artist
	Location
}

func main(){
	artist := Artist {
		ID: 1,
		Name: "Queen",
	}

	location := Location {
		ID: 1,
		Locations: []string {
			"England",
			"Germany",
		},
	}

	info := ArtistInfo {
		Artist: artist,
		Location: location,
	}

	fmt.Println(info.Name)

	for _, location := range info.Locations {
		fmt.Println(location)
	}
}