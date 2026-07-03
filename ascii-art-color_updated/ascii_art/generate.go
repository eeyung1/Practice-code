package ascii

import "strings"

func GenerateArt(
	text string,
	banner map[rune][]string,
	color string,
	substring string,
) string {

	var art strings.Builder

	if text == "" {
		return ""
	}

	if text == "\\n" {
		return "\n"
	}

	char, err := ValidateInput(text)
	if char != 0 && err != nil {
		return "Input contains none printable ascii character"
	}

	input := strings.Split(text, "\n")

	for _, line := range input {

		highlight := BuildHighlight(line, substring)

		if line != "" {
			art.WriteString(strings.Join(
				RenderLine(line, banner, color, highlight),
				"\n",
			))
			art.WriteString("\n")
		} else {
			art.WriteRune('\n')
		}
	}

	return art.String()
}
