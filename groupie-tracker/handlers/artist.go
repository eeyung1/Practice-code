package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/services"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(
			w,
			"400 Bad Request",
			http.StatusBadRequest,
		)

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
		template.ParseFiles("templates/artist.html"),
	)

	artist, found := services.FindArtistByID(
		id,
		artists,
	)

	if !found {
		http.Error(
			w,
			"404 Not Found",
			http.StatusNotFound,
		)

		return
	}

	tmpl.Execute(w, artist)
}
