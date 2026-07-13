package services

import (
	"groupie-tracker-geolocalization/models"
	"testing"
)

func TestFindArtistByIDFound(t *testing.T) {
	artists := []models.Artist {
		{
			ID: 1,
			Name: "Queen",
		},
		{
			ID: 2,
			Name: "ABBA",
		},
	}

	artist, found := FindArtistByID(
		2,
		artists,
	)

	if !found {
		t.Errorf(
			"expected artist to be found",
		)
	}

	if artist.Name != "ABBA" {
		t.Errorf(
			"expected ABBA, got %s",
			artist.Name,
		)
	}
}

func TestFindArtistByIDNotFound(t *testing.T) {
	artists := []models.Artist{
		{
			ID: 1,
			Name: "Queen",
		},
	}

	artist, found := FindArtistByID(
		99,
		artists,
	)

	if found {
		t.Errorf(
			"expected artist not to be found",
		)
	}

	if artist != nil {
		t.Errorf(
			"expected nil artist",
		)
	}
}