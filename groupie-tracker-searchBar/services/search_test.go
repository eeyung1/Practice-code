package services

import (
	"groupie-tracker-geolocalization/models"
	"testing"
)

func TestSearchArtists(t *testing.T) {
	artists := []models.Artist {
		{
			ID: 1,
			Name: "Queen",
		},
		{
			ID: 2,
			Name: "ABBA",
		},
		{
			ID: 3,
			Name: "Pink Floyd",
		},
	}

	tests := []struct{
		name 				string
		query 				string
		expectedCount		int
		expectedArtistName	string
	}{
		{
			name: "exact match",
			query: "Queen",
			expectedCount: 1,
			expectedArtistName: "Queen",
		},
		{
			name: "partial match",
			query: "ue",
			expectedCount: 1,
			expectedArtistName: "Queen",
		},
		{
			name: "case insensitive match",
			query: "QUEEN",
			expectedCount: 1,
			expectedArtistName: "Queen",
		},
		{
			name: "no match",
			query: "Metallica",
			expectedCount: 0,
			expectedArtistName: "",
		},
		{
			name: "empty query returns all artists",
			query: "",
			expectedCount: 3,
			expectedArtistName: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SearchArtists(
				test.query,
				artists,
			)

			if len(result) != test.expectedCount {
				t.Errorf(
					"expected %d results, got %d",
					test.expectedCount,
					len(result),
				)
			}

			if test.expectedArtistName != "" {
				if result[0].Name != test.expectedArtistName {
					t.Errorf(
						"expected %s, got %s",
						test.expectedArtistName,
						result[0].Name,
					)
				}
			}
		})
	}
}