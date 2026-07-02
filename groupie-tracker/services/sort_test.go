package services

import (
	"groupie-tracker/models"
	"testing"
)

func TestSortArtists(t *testing.T) {
	tests := []struct{
		name			string
		field			string
		order			string
		expectedOrder	[]string
	}{
		{
			name: "name ascending",
			field: "name",
			order: "asc",
			expectedOrder: []string{
				"ABBA",
				"Pink Floyd",
				"Queen",
			},
		},
		{
			name: "creation date descending",
			field: "creationDate",
			order: "desc",
			expectedOrder: []string{
				"Pink Floyd",
				"ABBA",
				"Queen",
			},
		},
		{
			name: "invalid field falls back to name ascending",
			field: "banana",
			order: "asc",
			expectedOrder: []string{
				"ABBA",
				"Pink Floyd",
				"Queen",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			artists := []models.Artist{
				{
					Name: "Queen",
					CreationDate: 1970,
				},
				{
					Name: "ABBA",
					CreationDate: 1973,
				},
				{
					Name: "Pink Floyd",
					CreationDate: 1991,
				},
			}

			SortArtists(
				artists,
				test.field,
				test.order,
			)

			for i, expectedName := range test.expectedOrder {
				if artists[i].Name != expectedName {
					t.Errorf(
						"expected %s at position %d, got %s",
						expectedName,
						i,
						artists[i].Name,
					)
				}
			}
		})
	}
}