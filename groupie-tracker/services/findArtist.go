package services

import "groupie-tracker/models"

func FindArtistByID(
	id int,
	artists []models.Artist,
)(*models.Artist, bool) {
	if id < 1 || id > len(artists) {
		return nil, false
	}

	return &artists[id-1], true
}