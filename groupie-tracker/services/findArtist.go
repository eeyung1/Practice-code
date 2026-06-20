package services

import "groupie-tracker/models"

func FindArtistByID(
	id int,
	artists []models.Artist,
)(*models.Artist, bool) {
	for _, artist := range artists {
		if artist.ID == id {
			return &artist, true
		}
	}

	return nil, false
}