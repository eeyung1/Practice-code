package services

import "groupie-tracker/models"

var artistCache []models.Artist

var relationCache = map[int]models.Relation{}