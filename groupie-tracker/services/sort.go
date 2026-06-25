package services

import (
	"groupie-tracker/models"
	"sort"
)

func SortArtists(
	artists []models.Artist,
	field string,
	order string,
) {
	sort.Slice(
		artists,
		func(i, j int) bool {
			if field == "creationDate" {
				if order == "desc" {
					return artists[i].CreationDate >
					artists[j].CreationDate
				}

				return artists[i].CreationDate <
				artists[j].CreationDate
			}

			if order == "desc" {
				return artists[i].Name >
				artists[j].Name
			}

			return artists[i].Name < artists[j].Name
		},
	)
}