package ascii

import "strings"

func RenderLine(
	text string,
	banner map[rune][]string,
	color string,
	highlight []bool,
) []string {

	container := make([]string, 8)

	for row := range 8 {

		var rowBuilder strings.Builder

		for i, char := range text {

			ascii := banner[char][row]

			if color != "" && i < len(highlight) && highlight[i] {
				rowBuilder.WriteString(color)
				rowBuilder.WriteString(ascii)
				rowBuilder.WriteString(Reset)
			} else {
				rowBuilder.WriteString(ascii)
			}
		}

		container[row] = rowBuilder.String()
	}

	return container
}
