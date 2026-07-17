package services

import (
	"encoding/json"
	"fmt"
	"groupie-tracker-geolocalization/models"
	"io"
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

var coordinateMap = make(
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

	// fmt.Println(
	// 	"Status:",
	// 	resp.Status,
	// )

	if err != nil {
		return models.Coordinate{}, err
	}

	defer resp.Body.Close()

	var response []nominatimResponse

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Coordinate{}, err
	}


	err = json.Unmarshal(body, &response)

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

	coordinate := models.Coordinate{
		Latitude:  latitude,
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
			fmt.Println(
				"Could not geocode:",
				location,
				err,
			)

			continue
		}

		coordinateMap[location] = coordinate
	}

	return coordinateMap, nil
}

func InitializeCoordinates(
    relations []models.Relation,
) error {

    locations := GetUniqueLocations(relations)

    coordinates, err := GeocodeLocations(locations)
    if err != nil {
        return err
    }

    coordinateMap = coordinates

    return nil
}
