package main

import (
	"strings"
)

func Render(input string, banner map[rune][]string) string {
	lines := strings.Split(input, "\\n")

	var result strings.Builder

	for lineIndex, line := range lines {
		if line == "" {
			if lineIndex < len(lines)-1 {
				result.WriteString("\n")
			}
			continue
		}

		for i := 0; i < 8; i++ {
			for _, ch := range line {
				art, ok := banner[ch]
				if !ok {
					art = banner[' ']
				}

				result.WriteString(art[i])
			}

			result.WriteString("\n")
		}
	}

	return result.String()
}
