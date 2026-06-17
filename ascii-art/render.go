package main

import "strings"

func RenderLine(input string, banner map[rune][]string) []string {
	output := make([]string, 8)

	for index := range 8 {
		var result strings.Builder

		for _, char := range input {
			art, ok := banner[char]
			if !ok {
				art = banner[' ']
			}

			result.WriteString(art[index])
		}

		output[index] = result.String()
	}

	return output

}
