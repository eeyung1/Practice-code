package main

import "strings"

func PadArtRow(rows []string, width int) []string {
	result := make([]string, len(rows))

	for i, row := range rows {
		if width > len(row) {
			result[i] = row + strings.Repeat(" ", width-len(row))
		} else {
			result[i] = row
		}
	}
	return result
}
