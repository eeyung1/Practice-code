package main

import "strings"

func TrimArtRows(rows []string) []string {
	result := make([]string, len(rows))
	for i, v := range rows {
		if strings.Trim(v, " ") == " " {
			result[i] = strings.TrimRight(v, " ")
		}

	}
	return result

}
