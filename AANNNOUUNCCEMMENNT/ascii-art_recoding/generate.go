package main

import "strings"

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}

	ch, err := ValidateInput(input)
	if ch != 0 && err != nil {
		return "input has unsupported characters"
	}

	parts := SplitInput(input)

	var result strings.Builder

	onlyNewLines := true
	for _, part := range parts {
		if part != "" {
			onlyNewLines = false
			break
		}
	}

	if onlyNewLines {
		for i := 0; i < len(parts)-1; i++ {
			result.WriteByte('\n')
		}
		return result.String()
	}

	// Normal rendering
	for _, part := range parts {
		if part == "" {
			result.WriteByte('\n') // preserve blank line
			continue
		}

		rows := RenderLine(part, banner)
		for _, row := range rows {
			result.WriteString(row)
			result.WriteByte('\n')
		}
	}

	return result.String()
}