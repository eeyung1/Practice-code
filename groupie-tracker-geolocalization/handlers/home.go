package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"groupie-tracker-geolocalization/models"
	"groupie-tracker-geolocalization/services"
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
			"500 Internal Server 1Error",
			http.StatusInternalServerError,
		)

		return
	}

	relations, err := services.GetRelations()

	if err != nil {
		http.Error(
			w,
			"500 Internal Server 2Error",
			http.StatusInternalServerError,
		)
		return
	}

	relationMap := services.BuildRelationMap(relations)

	coordinateMap := services.GetCoordinateMap()

	countries := services.GetUniqueCountries(relations)

	filters := services.ParseFilters(r)

	artists = services.FilterByCreationDate(
		artists,
		filters.CreationMin,
		filters.CreationMax,
	)

	artists = services.FilterByFirstAlbum(
		artists,
		filters.AlbumMin,
		filters.AlbumMax,
	)

	if len(filters.Members) > 0 {
		artists = services.FilterByMembers(artists, filters.Members)
	}

	if len(filters.Countries) > 0 {
		artists = services.FilterByLocation(artists, relationMap, filters.Countries)
	}

	if filters.Query != "" {
		artists = services.SearchArtists(
			filters.Query,
			artists,
		)
	}

	services.SortArtists(
		artists,
		filters.Field,
		filters.Order,
	)

	tmpl := template.Must(
		template.New("index.html").
			Funcs(template.FuncMap{
				"containsMember":  ContainsMember,
				"containsCountry": ContainsCountry,
			}).
			ParseFiles("templates/index.html"),
	)

	coordinatesJSON, err := json.Marshal(coordinateMap)

	if err != nil {
		http.Error(
			w,
			"500 Internal Server 4Error",
			http.StatusInternalServerError,
		)

		return
	}

	data := models.PageData{
		Artists: artists,

		Query: filters.Query,
		Order: filters.Order,
		Field: filters.Field,

		CreationMin: fmt.Sprintf("%d", filters.CreationMin),
		CreationMax: fmt.Sprintf("%d", filters.CreationMax),

		AlbumMin: fmt.Sprintf("%d", filters.AlbumMin),
		AlbumMax: fmt.Sprintf("%d", filters.AlbumMax),

		SelectedMembers: filters.Members,

		Countries:         countries,
		SelectedCountries: filters.Countries,

		Coordinates: coordinateMap,

		CoordinatesJSON: string(coordinatesJSON),
	}

	// if err := tmpl.Execute(w, data); err != nil {
	// 	fmt.Println(err)
	// }

	if err := tmpl.Execute(w, data); err != nil {
		fmt.Println("Template error:", err)
		http.Error(
			w,
			"500 Internal Server Error",
			http.StatusInternalServerError,
		)
		return
	}
}
