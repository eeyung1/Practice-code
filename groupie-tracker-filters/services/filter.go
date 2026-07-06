package services

import (
	"groupie-tracker-filters/models"
	"strconv"
)

func FilterByCreationDate(
	artists []models.Artist,
	min int,
	max int,
) []models.Artist {
	var filtered []models.Artist

	for _, artist := range artists {
		if artist.CreationDate >= min && artist.CreationDate <= max {
			filtered = append(filtered, artist)
		}
	}

	return filtered
}

func FilterByFirstAlbum(
	artists []models.Artist,

	min int,
	max int,
) []models.Artist {
	
	var filtered []models.Artist

	for _, artist := range artists {
		yearStr := artist.FirstAlbum[len(artist.FirstAlbum)-4:]

		year, err := strconv.Atoi(yearStr)

		if err != nil {
			continue
		}

		if year >= min && year <= max {
			filtered = append(filtered, artist)
		}
	}

	return filtered
}