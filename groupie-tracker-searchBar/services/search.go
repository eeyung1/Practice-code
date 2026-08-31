package services

import (
	"groupie-tracker-geolocalization/models"
	"strconv"
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

func BuildSearchIndex(
	artists []models.Artist,
	relations []models.Relation,
) []models.SearchItem {

	searchIndex := make([]models.SearchItem, 0)

	artistByID := make(map[int]models.Artist)

	for _, artist := range artists {
		artistByID[artist.ID] = artist
	}

	for _, relation := range relations {
		artist, exists := artistByID[relation.ID]

		if !exists {
			continue
		}

		searchIndex = append(
			searchIndex,
			models.SearchItem{
				Value:    artist.Name,
				Type:     "artist",
				ArtistID: artist.ID,
			},
		)

		for _, member := range artist.Members {
			searchIndex = append(
				searchIndex,
				models.SearchItem{
					Value:    member,
					Type:     "member",
					ArtistID: artist.ID,
				},
			)
		}

		searchIndex = append(
			searchIndex,
			models.SearchItem{
				Value:    strconv.Itoa(artist.CreationDate),
				Type:     "creationDate",
				ArtistID: artist.ID,
			},
		)

		searchIndex = append(
			searchIndex,
			models.SearchItem{
				Value:    artist.FirstAlbum,
				Type:     "firstAlbum",
				ArtistID: artist.ID,
			},
		)
	}

	return searchIndex
}
