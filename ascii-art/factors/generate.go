package factors

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

	// Check if the input is exclusively newlines
	onlyNewLines := true
	for _, part := range parts {
		if part != "" {
			onlyNewLines = false
			break
		}
	}

	var result strings.Builder

	for i, part := range parts {
		if part == "" {
			// If it's only newlines, standard behavior drops the final trailing split element
			if onlyNewLines && i == len(parts)-1 {
				continue
			}
			result.WriteByte('\n')
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
