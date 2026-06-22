package services

import (
	"groupie-tracker/models"
	"sort"
)

func SortArtists(
	artists []models.Artist,
	order string,
) {
	sort.Slice(
		artists,
		func(i, j int) bool {
			if order == "desc" {
				return artists[i].Name > artists[j].Name
			}

			return artists[i].Name < artists[j].Name
		},
	)
}