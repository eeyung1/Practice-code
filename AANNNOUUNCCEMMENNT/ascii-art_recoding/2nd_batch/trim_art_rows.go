package main

import "strings"

func TrimArtRows(rows []string) []string {
	result := make([]string, len(rows))

	for i, row := range rows {
		result[i] = strings.TrimRight(row, " ")
	}

	return result
}