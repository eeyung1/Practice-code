package factors

import "strings"

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}
	if input == "\\n" {
		return "\n"
	}
	words := SplitInput(input)
	var output strings.Builder
	for _, word := range words {
		if word == "" {
			output.WriteString("\n")
			continue
		}
		row := RenderLine(word, banner)
		for _, rows := range row {
			output.WriteString(rows)
			output.WriteString("\n")
		}

	}
	return output.String()

}
