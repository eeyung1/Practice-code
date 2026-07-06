package services

import (
	"encoding/json"
	"groupie-tracker-filters/models"
	"net/http"
	"strconv"
)

var relationBaseURL = "https://groupietrackers.herokuapp.com/api/relation/"

func GetRelation(
	id int,
) (models.Relation, error) {

	if relation, exists := relationCache[id]; exists {
		return relation, nil
	}
	resp, err := http.Get(
		relationBaseURL + strconv.Itoa(id),
	)

	if err != nil {
		return models.Relation{}, err
	}

	defer resp.Body.Close()

	var relation models.Relation

	err = json.NewDecoder(resp.Body).Decode(&relation)

	if err != nil {
		return models.Relation{}, err
	}

	relationCache[id] = relation

	return relation, nil
}