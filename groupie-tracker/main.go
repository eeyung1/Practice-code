package main

import (
	"fmt"
	"groupie-tracker/models"
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	tmpl := template.Must(
		template.ParseFiles("templates/index.html"),
	)

	artists := []models.Artist{
		{
			ID: 1,
			Name: "Queen",
		},
		{
			ID: 2,
			Name: "ABBA",
		},
		{
			ID: 3,
			Name: "Pink Floyd",
		},
	}


	tmpl.Execute(w, artists)

}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Artist Page")
}

func main(){
	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/artist", artistHandler)

	fmt.Println("Server started on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}