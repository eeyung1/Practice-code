package services

import "groupie-tracker-filters/models"

var artistCache []models.Artist

var relationCache = map[int]models.Relation{}