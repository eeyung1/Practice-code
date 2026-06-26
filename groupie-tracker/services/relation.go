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
	resp, err := http.Get(
		"https://groupietrackers.herokuapp.com/api/relation/" + 
		strconv.Itoa(id),
	)

	if err != nil {
		return models.Relation{}, nil
	}

	defer resp.Body.Close()

	var relation models.Relation

	err = json.NewDecoder(resp.Body).Decode(&relation)

	if err != nil {
		return models.Relation{}, nil
	}

	return relation, nil
}