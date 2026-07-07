package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"groupie-tracker-filters/models"
	"groupie-tracker-filters/services"
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

	creationMin := r.URL.Query().Get("creationMin")
	creationMax := r.URL.Query().Get("creationMax")

	albumMin := r.URL.Query().Get("albumMin")
	albumMax := r.URL.Query().Get("albumMax")

	memberValues := r.URL.Query()["members"]

	var selectedMembers []int

	for _, value := range memberValues {
		memberCount, err := strconv.Atoi(value)
		if err != nil {
			continue
		}

		selectedMembers = append(selectedMembers, memberCount)
	}

	albumMinYear := 0
	albumMaxYear := 9999

	if albumMin != "" {
		albumMinYear, _ = strconv.Atoi(albumMin)
	}

	if albumMax != "" {
		albumMaxYear, _ = strconv.Atoi(albumMax)
	}

	min := 0
	max := 9999

	if creationMin != "" {
		min, _ = strconv.Atoi(creationMin)
	}

	if creationMax != "" {
		max, _ = strconv.Atoi(creationMax)
	}

	artists = services.FilterByCreationDate(
		artists,
		min,
		max,
	)

	artists = services.FilterByFirstAlbum(
		artists,
		albumMinYear,
		albumMaxYear,
	)

	if len(selectedMembers) > 0 {
		artists = services.FilterByMembers(artists, selectedMembers)
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

	// tmpl := template.Must(
	// 	template.ParseFiles("templates/index.html"),
	// )

	tmpl := template.Must(
		template.New("index.html").Funcs(template.FuncMap{
			"containsMember": ContainsMember,
		}).ParseFiles("templates/index.html"),
	)

	data := models.PageData{
		Artists: artists,

		Query: query,
		Order: order,
		Field: field,

		CreationMin: creationMin,
		CreationMax: creationMax,

		AlbumMin: albumMin,
		AlbumMax: albumMax,

		SelectedMembers: selectedMembers,
	}

	if err := tmpl.Execute(w, data); err != nil {
		fmt.Println(err)
	}
}
