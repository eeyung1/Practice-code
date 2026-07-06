package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"groupie-tracker-filters/models"
	"groupie-tracker-filters/services"
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

	var (
		artists  []models.Artist
		relation models.Relation

		artistErr   error
		relationErr error
	)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		artists, artistErr = services.GetArtists()
	}()

	go func() {
		defer wg.Done()

		relation, relationErr = services.GetRelation(id)

	}()

	wg.Wait()

	if artistErr != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	if relationErr != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
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

	data := models.ArtistPageData{
		Artist:   *artist,
		Relation: relation,
	}

	if err := tmpl.Execute(w, data); err != nil {
		fmt.Println(err)
	}
}
