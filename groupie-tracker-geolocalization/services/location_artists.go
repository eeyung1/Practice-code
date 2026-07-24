package services

import "groupie-tracker-geolocalization/models"

func BuildLocationArtists(
	artists []models.Artist,
	relations []models.Relation,
) map[string][]string {

	locationArtists := make(map[string][]string)

	artistMap := make(map[int]string)

	for _, artist := range artists {
		artistMap[artist.ID] = artist.Name
	}

	for _, relation := range relations {

		artistName := artistMap[relation.ID]

		for location := range relation.DatesLocations {
			locationArtists[location] = append(
				locationArtists[location],
				artistName,
			)
		}
	}

	return locationArtists
}
