package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 - Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "400 - Bad Request: method not allowed", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Home page. Method: %s", r.Method)
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "400 - Bad Request: method not allowed", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "ASCII art page. Method: %s", asciiArtHandler)
}

func main(){
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}