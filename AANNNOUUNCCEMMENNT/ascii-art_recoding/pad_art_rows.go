package main

import (
	"strings"
)

func PadArtRows(rows []string, width int) []string {
	result := make([]string, len(rows))

	for i, v := range rows {
		if width <= 0 {
			result[i] = v
		}

		padding := width - len(v)

		if padding > 0 {
			result[i] = v + strings.Repeat(" ", padding)
		} else {
			result[i] = v
		}
	}

	return result
}