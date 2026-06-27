package main

import (
	"unicode"
)

func countFrequencies(text string) map[rune]int {
	count := make(map[rune]int)

	for _, ch := range text {
		lower := unicode.ToLower(ch)

		if lower >= 'a' && lower <= 'z' {
			count[lower]++
		}
	}

	return count
}