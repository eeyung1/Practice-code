package services

import "groupie-tracker-geolocalization/models"

var artistCache []models.Artist

var relationCache = map[int]models.Relation{}