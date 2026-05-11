package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 - Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "404 - Bad Request: method not allowed", http.StatusBadRequest)
		return
	}

	html, err := os.ReadFile("form.html")
	if err != nil {
		http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(html)
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "400 - Bad Request: method not allowed", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "400 - Bad Request: could not parse form", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	fmt.Println("--- Form submission received ---")
	fmt.Println("text:", text)
	fmt.Println("banner:", banner)

	fmt.Fprintf(w, "Received!\nText: %s\nBanner: %s", text, banner)
}

func main(){
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}