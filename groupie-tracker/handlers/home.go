package handlers

import (
	"html/template"
	"net/http"

	"groupie-tracker/services"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	artists, err := services.GetArtists()

	if err != nil {
		http.Error(
			w,
			"500 Internal Server Error",
			http.StatusInternalServerError,
		)

		return
	}

	tmpl := template.Must(
		template.ParseFiles("templates/index.html"),
	)

	tmpl.Execute(w, artists)
}