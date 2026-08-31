package services

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetArtists(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					t.Errorf(
						"expected Get request, got %s",
						r.Method,
					)
				}

				w.Header().Set("Content-Type", "application/json")

				w.Write([]byte(`
				[
					{
						"id": 1,
						"image": "queen.jpg",
						"name": "Queen",
						"members": ["Freddie Mercury"],
						"creationDate": 1970,
						"firstAlbum": "1973"
					}
				]
				
				`))
			},

			
		),
	)

	defer server.Close()

	originalURL := artistsURL

	artistsURL = server.URL

	defer func ()  {
		artistsURL = originalURL
	}()

	artists, err := GetArtists()

	if err != nil {
		t.Fatalf(
			"expected no error, got %v",
			err,
		)
	}

	if len(artists) != 1 {
		t.Fatalf(
			"expected 1 artist, got %d",
			len(artists),
		)
	}

	if artists[0].Name != "Queen" {
		t.Errorf(
			"expected Queen, got %s",
			artists[0].Name,
		)
	}

	if artists[0].CreationDate != 1970 {
		t.Errorf(
			"expected creation date 1970, got %d",
			artists[0].CreationDate,
		)
	}
}