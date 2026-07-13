package services

import (
	"encoding/json"
	"groupie-tracker-geolocalization/models"
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

func GetRelations()([]models.Relation, error) {
	resp, err := http.Get(
		"https://groupietrackers.herokuapp.com/api/relation",
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response struct {
		Index []models.Relation	`json:"index"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Index, nil
}

func BuildRelationMap(
	relations []models.Relation,
) map[int]models.Relation {
	relationMap := make(map[int]models.Relation)

	for _, relation := range relations {
		relationMap[relation.ID] = relation
	}

	return relationMap
}