package services

import (
	"encoding/json"
	"io"
	"net/http"

	"groupie-tracker/models"
)

func GetArtists() ([]models.Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

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

	return artists, nil
}


