package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii-art-web/handlers"
)

func main() {
	// Register route handlers
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiArtHandler)

	port := ":8080"
	fmt.Println("Server running at http://localhost" + port)

	// Start the server; log.Fatal will print and exit if it fails
	log.Fatal(http.ListenAndServe(port, nil))
}
