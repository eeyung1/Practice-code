package services

import (
	"groupie-tracker-geolocalization/models"
	"sort"
	"strconv"
	"strings"
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

func FilterByMembers(
	artists []models.Artist,
	selected []int,
) []models.Artist {
	var filtered []models.Artist

	for _, artist := range artists {
		
		memberCount := len(artist.Members)

		for _, choice := range selected {
			if memberCount == choice {
				filtered = append(filtered, artist)
				break
			}
		}
	}

	return filtered
}


func GetUniqueCountries(
	relations	[]models.Relation,
) [] string {
	countrySet := make(map[string]struct{})

	for _, relation := range relations {
		for location := range relation.DatesLocations {
			parts := strings.Split(location, "-")

			country := parts[len(parts)-1]

			countrySet[country] = struct{}{}
		}
	}

	var countries []string

	for country := range countrySet {
		countries = append(countries, country)
	}

	sort.Strings(countries)

	return countries
}

func GetUniqueLocations(
    relations []models.Relation,
) []string {
	locationSet := make(map[string]struct{})

	for _, relation := range relations {
		for location := range relation.DatesLocations {
			locationSet[location] = struct{}{}
		}
	}

	var locations []string

	for location := range locationSet {
		locations = append(locations, location)
	}

	sort.Strings(locations)

	return locations
}

func FilterByLocation(
	artists []models.Artist,
	relationMap map[int]models.Relation,
	selectedCountries []string,
) []models.Artist {

	var filtered []models.Artist

	for _, artist := range artists {

		relation := relationMap[artist.ID]

		found := false

		for location := range relation.DatesLocations {

			parts := strings.Split(location, "-")
			country := parts[len(parts)-1]

			for _, selected := range selectedCountries {

				if strings.EqualFold(country, selected) {
					filtered = append(filtered, artist)
					found = true
					break
				}
			}

			if found {
				break
			}
		}
	}

	return filtered
}