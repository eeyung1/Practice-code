package main

import (
	"fmt"
	"groupie-tracker/handlers"

	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)

	http.HandleFunc("/artist/", handlers.ArtistHandler)

	fmt.Println("Server started on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
