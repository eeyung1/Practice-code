package main

import (
	"fmt"
	"groupie-tracker-geolocalization/handlers"
	"groupie-tracker-geolocalization/services"

	"net/http"
)

func main() {

	relations, err := services.GetRelations()

	if err != nil {
		panic(err)
	}

	err = services.InitializeCoordinates(relations)

	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", handlers.HomeHandler)

	http.HandleFunc("/artist/", handlers.ArtistHandler)

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server started on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
