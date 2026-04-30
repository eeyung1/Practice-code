package main

import (
	// "fmt"
	"strings"
)

func generateArt(text string, banner map[rune][]string) string {
	if len(text) == 0 {
		return ""
	}
	ch, err := ValidateInput(text)
	if ch != 0 && err != nil {
		return "input has unsupported characters"
	}
	segments := SplitInput(text)
	onlyNewLines := true

	var result strings.Builder
	for _, ch := range segments {
		if ch != "" {
			onlyNewLines = false
			break
		}
	}
	if onlyNewLines {
		for i := 0; i < len(segments)-1; i++ {
			result.WriteByte('\n')
		}
	} else {
		for _, ch := range segments {
			if ch == "" {
				result.WriteByte('\n')
				continue
			}
			newText := strings.Join(RenderLine(ch, banner), "\n")
			result.WriteString(newText)
		}
	}
	return result.String()
}


func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}

	parts := SplitInput(input)

	onlyNewLines := true

	for _, part := range parts {
		if part != "" {
			onlyNewLines = false
			break
		}
	}

	var result strings.Builder

	if onlyNewLines {
		for i := 0; i < len(parts)-1; i++ {
			result.WriteByte('\n')
		}

		return result.String()
	}

	for i, part := range parts {
		if part == "" {
			if i < len(parts) - 1 {
				result.WriteString("\n")
			}
		} else {
			rows := RenderLine(part, banner)
			for _, row := range rows {
				result.WriteString(row + "\n")
			}
		}
	}
	return result.String()
}
