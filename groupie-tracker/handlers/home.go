package handlers

import (
	"html/template"
	"net/http"

	"groupie-tracker/models"
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

	query := r.URL.Query().Get("search")
	field := r.URL.Query().Get("field")

	if query != "" {
		artists = services.SearchArtists(
			query,
			artists,
		)
	}

	order := r.URL.Query().Get("sort")

	services.SortArtists(
		artists,
		field,
		order,
	)

	tmpl := template.Must(
		template.ParseFiles("templates/index.html"),
	)

	data := models.PageData{
		Artists: artists,
		Query: query,
		Order: order,
	}

	tmpl.Execute(w, data)
}