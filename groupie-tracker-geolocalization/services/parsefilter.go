package services

import (
	"groupie-tracker-geolocalization/models"
	"net/http"
	"strconv"
)

func ParseFilters(r *http.Request) models.FilterOptions {
	query := r.URL.Query().Get("search")
	order := r.URL.Query().Get("sort")

	field := r.URL.Query().Get("field")
	creationMin := r.URL.Query().Get("creationMin")
	creationMax := r.URL.Query().Get("creationMax")

	albumMin := r.URL.Query().Get("albumMin")
	albumMax := r.URL.Query().Get("albumMax")

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

	memberValues := r.URL.Query()["members"]

	var selectedMembers []int

	for _, value := range memberValues {
		memberCount, err := strconv.Atoi(value)
		if err != nil {
			continue
		}

		selectedMembers = append(selectedMembers, memberCount)
	}

	selectedCountries := r.URL.Query()["countries"]

	final := models.FilterOptions{
		Query: query,
		Order: order,
		Field: field,

		CreationMin: min,
		CreationMax: max,

		AlbumMin: albumMinYear,
		AlbumMax: albumMaxYear,

		Members: selectedMembers,

		Countries: selectedCountries,
	}

	return final
}
