package services

import (
	"encoding/json"
	"io"
	"net/http"

	"groupie-tracker/models"
)

var artistsURL = "https://groupietrackers.herokuapp.com/api/artists"

func GetArtists() ([]models.Artist, error) {

	if len(artistCache) > 0 {
		return artistCache, nil
	}
	resp, err := http.Get(artistsURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var artists []models.Artist

	err = json.Unmarshal(body, &artists)

	if err != nil {
		return nil, err
	}

	artistCache = artists

	return artists, nil
}


