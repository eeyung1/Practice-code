package main

import (
	"fmt"
	"groupie-tracker/services"

	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	artists, err := services.GetArtists()
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(
		template.ParseFiles("templates/index.html"),
	)
	tmpl.Execute(w, artists)

}

func main() {
	http.HandleFunc("/", homeHandler)

	// http.HandleFunc("/artist", artistHandler)

	fmt.Println("Server started on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
