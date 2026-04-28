package main

import (
	"strings"
)

func RenderLine(input string, banner map[rune][]string) []string {
	output := make([]string, 8)

	for i := range 8 {
		var builder strings.Builder

		for _, ch := range input {
			art, ok := banner[ch]

			if !ok {
				art = banner[' ']
			}

			builder.WriteString(art[i])
		}

		output[i] = builder.String()
	}

	return output
}


