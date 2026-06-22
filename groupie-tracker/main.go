package main

import (
	"fmt"
	"groupie-tracker/handlers"

	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)

	http.HandleFunc("/artist/", handlers.ArtistHandler)

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server started on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
