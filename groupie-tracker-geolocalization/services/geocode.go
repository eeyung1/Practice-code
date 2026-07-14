package services

import (
	"encoding/json"
	"fmt"
	"groupie-tracker-geolocalization/models"
	"net/http"
	"net/url"
	"strconv"
)

type nominatimResponse struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

var coordinateCache = make(
	map[string]models.Coordinate,
)

func GeocodeLocation(
	location string,
) (models.Coordinate, error) {

	if coordinate, exists := coordinateCache[location]; exists {
		return coordinate, nil
	}

	escapedLocation := url.QueryEscape(location)

	requestURL := fmt.Sprintf(
		"https://nominatim.openstreetmap.org/search?q=%s&format=json&limit=1",
		escapedLocation,
	)

	req, err := http.NewRequest(
		http.MethodGet,
		requestURL,
		nil,
	)

	if err != nil {
		return models.Coordinate{}, err
	}

	req.Header.Set(
		"User-Agent",
		"GroupieTrackerGeolocalization/1.0 (Learning Project)",
	)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return models.Coordinate{}, err
	}

	defer resp.Body.Close()

	var response []nominatimResponse

	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		return models.Coordinate{}, err
	}

	if len(response) == 0 {
		return models.Coordinate{}, fmt.Errorf("location not found")
	}

	latitude, err := strconv.ParseFloat(response[0].Lat, 64)

	if err != nil {
		return models.Coordinate{}, err
	}

	longitude, err := strconv.ParseFloat(response[0].Lon, 64)

	if err != nil {
		return models.Coordinate{}, err
	}

	coordinate := models.Coordinate {
		Latitude: latitude,
		Longitude: longitude,
	}

	coordinateCache[location] = coordinate

	return coordinate, nil
}

func GeocodeLocations(
	locations []string,
) (map[string]models.Coordinate, error) {
	coordinateMap := make(map[string]models.Coordinate)

	for _, location := range locations {
		coordinate, err := GeocodeLocation(location)

		if err != nil {
			return nil, err
		}

		coordinateMap[location] = coordinate
	}

	return coordinateMap, nil
}
