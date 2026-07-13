package handlers

import "strings"

func ContainsMember(
	selected []int,
	member int,
) bool {
	for _, value := range selected {
		if value == member {
			return true
		}
	}

	return false
}

func ContainsCountry(
	selected []string,
	country string,
) bool {

	for _, value := range selected {
		if strings.EqualFold(value, country) {
			return true
		}
	}

	return false
}
