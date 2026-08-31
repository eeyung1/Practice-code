package services

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRelation(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					t.Errorf(
						"expected GET request, got %s",
						r.Method,
					)
				}

				w.Header().Set("Content-Type", "application/json")

				w.Write([]byte(`
				{
					"id": 1,
					"datesLocations": {
						"london-uk": [
							"01-01-2020",
							"02-01-2020"
						]
					}
				}

				`))
			},
		),
	)

	defer server.Close()

	originalURL := relationBaseURL

	relationBaseURL = server.URL + "/"

	defer func ()  {
		relationBaseURL = originalURL
	}()

	relation, err := GetRelation(1)

	if err != nil {
		t.Fatalf(
			"expected no error, got %v",
			err,
		)
	}

	if relation.ID != 1 {
		t.Errorf(
			"expected ID 1, got %d",
			relation.ID,
		)
	}

	dates, exists := relation.DatesLocations["london-uk"]

	if !exists {
		t.Fatalf(
			"expected london-uk location to exist",
		)
	}

	if len(dates) != 2 {
		t.Errorf(
			"expected 2 dates, got %d",
			len(dates),
		)
	}

	if dates[0] != "01-01-2020" {
		t.Errorf(
			"expected first date to be 01-01-2020, got %s",
			dates[0],
		)
	}
}