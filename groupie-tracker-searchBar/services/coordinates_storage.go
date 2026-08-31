package services

import (
	"encoding/json"
	"os"

	"groupie-tracker-geolocalization/models"
)

func SaveCoordinates(
	coordinates map[string]models.Coordinate,
) error {

	file, err := os.Create("coordinates.json")
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	return encoder.Encode(coordinates)
}

func LoadCoordinates() (
	map[string]models.Coordinate,
	error,
) {
	file, err := os.Open("coordinates.json")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	coordinates := make(map[string]models.Coordinate)

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&coordinates)
	if err != nil {
		return nil, err
	}

	return coordinates, nil
}