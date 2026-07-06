package services

import (
	"groupie-tracker-filters/models"
	"strings"
)

func SearchArtists(
	query string,
	artists []models.Artist,
) []models.Artist {
	var result []models.Artist

	query = strings.ToLower(query)

	for _, artist := range artists {
		if strings.Contains(
			strings.ToLower(artist.Name),
			query,
		) {
			result = append(result, artist)
		}
	}

	return result
}