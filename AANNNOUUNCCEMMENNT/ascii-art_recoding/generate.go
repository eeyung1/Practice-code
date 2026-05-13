package main

import "strings"

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}

	parts := SplitInput(input)

	var result strings.Builder

	// Normal rendering
	for i, part := range parts {
		if part == "" {
			if i < len(parts)-1 {
				result.WriteByte('\n')
			} // preserve blank line
			continue
		}

		rows := RenderLine(part, banner)
		for _, row := range rows {
			result.WriteString(row)
			result.WriteByte('\n')
		}

		if i == len(parts)-2 && parts[i+1] == "" {
			for j := 0; j < 8; j++ {
				result.WriteString("\n")
			}

			break
		}
	}

	return result.String()
}
