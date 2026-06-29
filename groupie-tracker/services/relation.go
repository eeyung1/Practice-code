package services

import (
	"encoding/json"
	"groupie-tracker/models"
	"net/http"
	"strconv"
)

func GetRelation(
	id int,
) (models.Relation, error) {

	if relation, exists := relationCache[id]; exists {
		return relation, nil
	}
	resp, err := http.Get(
		"https://groupietrackers.herokuapp.com/api/relation/" + 
		strconv.Itoa(id),
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