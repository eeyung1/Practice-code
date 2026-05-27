package builder

import "strings"

type segment struct {
	text  string
	style string
}

type ArtBuilder struct {
	segments []segment
}

func NewArtBuilder() *ArtBuilder {
	return &ArtBuilder{
		segments: []segment{},
	}
}

func (ab *ArtBuilder) AddText(text string) *ArtBuilder {
	ab.segments = append(ab.segments, segment{text: text, style: "normal"})
	return ab
}

func (ab *ArtBuilder) SetStyle(style string) *ArtBuilder {
	validStyles := map[string]bool{
		"normal": true, "bold": true,
		"italic": true, "outline": true,
	}
	if !validStyles[style] {
		panic("unsupported style: " + style)
	}
	if len(ab.segments) > 0 {
		ab.segments[len(ab.segments)-1].style = style
	}
	return ab
}

func basePattern(c rune) []string {
	patterns := map[rune][]string{
		'A': {"  **  ", " *  * ", " *  * ", " **** ", " *  * ", " *  * ", " *  * ", "      "},
		'B': {" ***  ", " *  * ", " *  * ", " ***  ", " *  * ", " *  * ", " ***  ", "      "},
		'C': {"  *** ", " *    ", " *    ", " *    ", " *    ", " *    ", "  *** ", "      "},
		'D': {" ***  ", " *  * ", " *  * ", " *  * ", " *  * ", " *  * ", " ***  ", "      "},
		'E': {" **** ", " *    ", " *    ", " ***  ", " *    ", " *    ", " **** ", "      "},
		'F': {" **** ", " *    ", " *    ", " ***  ", " *    ", " *    ", " *    ", "      "},
		'G': {"  *** ", " *    ", " *    ", " * ** ", " *  * ", " *  * ", "  *** ", "      "},
		'H': {" *  * ", " *  * ", " *  * ", " **** ", " *  * ", " *  * ", " *  * ", "      "},
		'I': {" **** ", "  **  ", "  **  ", "  **  ", "  **  ", "  **  ", " **** ", "      "},
		'J': {"   ** ", "   *  ", "   *  ", "   *  ", "   *  ", " * *  ", "  **  ", "      "},
		'K': {" *  * ", " * *  ", " **   ", " **   ", " * *  ", " * *  ", " *  * ", "      "},
		'L': {" *    ", " *    ", " *    ", " *    ", " *    ", " *    ", " **** ", "      "},
		'M': {" *  * ", " ** * ", " * ** ", " *  * ", " *  * ", " *  * ", " *  * ", "      "},
		'N': {" *  * ", " ** * ", " ** * ", " * ** ", " * ** ", " *  * ", " *  * ", "      "},
		'O': {"  **  ", " *  * ", " *  * ", " *  * ", " *  * ", " *  * ", "  **  ", "      "},
		'P': {" ***  ", " *  * ", " *  * ", " ***  ", " *    ", " *    ", " *    ", "      "},
		'Q': {"  **  ", " *  * ", " *  * ", " *  * ", " * ** ", " *  * ", "  *** ", "      "},
		'R': {" ***  ", " *  * ", " *  * ", " ***  ", " **   ", " * *  ", " *  * ", "      "},
		'S': {"  *** ", " *    ", " *    ", "  **  ", "    * ", "    * ", " ***  ", "      "},
		'T': {" **** ", "  **  ", "  **  ", "  **  ", "  **  ", "  **  ", "  **  ", "      "},
		'U': {" *  * ", " *  * ", " *  * ", " *  * ", " *  * ", " *  * ", "  **  ", "      "},
		'V': {" *  * ", " *  * ", " *  * ", " *  * ", " *  * ", "  **  ", "  **  ", "      "},
		'W': {" *  * ", " *  * ", " *  * ", " *  * ", " * ** ", " ** * ", " *  * ", "      "},
		'X': {" *  * ", " *  * ", "  **  ", "  **  ", "  **  ", " *  * ", " *  * ", "      "},
		'Y': {" *  * ", " *  * ", " *  * ", "  **  ", "  **  ", "  **  ", "  **  ", "      "},
		'Z': {" **** ", "    * ", "   *  ", "  *   ", " *    ", " *    ", " **** ", "      "},
	}

	if p, ok := patterns[c]; ok {
		return p
	}

	// fallback for any character not in the map
	row := " " + string(c) + "    "
	return []string{row, row, row, row, row, row, row, "      "}
}

func applyNormal(rows []string) []string {
	return rows
}

func applyBold(rows []string) []string {
	result := make([]string, len(rows))
	for i, row := range rows {
		var bold strings.Builder
		for _, ch := range row {
			bold.WriteRune(ch)
			bold.WriteRune(ch) // every character written twice
		}
		result[i] = bold.String()
	}
	return result
}

func applyItalic(rows []string) []string {
	result := make([]string, len(rows))
	total := len(rows)
	for i, row := range rows {
		// top rows shift right more, bottom rows shift less
		shift := total - 1 - i
		result[i] = strings.Repeat(" ", shift) + row
	}
	return result
}

func applyOutline(rows []string) []string {
	if len(rows) == 0 {
		return rows
	}
	width := len(rows[0])
	border := "+" + strings.Repeat("-", width) + "+"
	result := make([]string, 0, len(rows)+2)
	result = append(result, border)
	for _, row := range rows {
		result = append(result, "|"+row+"|")
	}
	result = append(result, border)
	// outline adds 2 extra rows — trim last 2 base rows to keep total at 8
	if len(result) > 8 {
		result = result[:8]
	}
	return result
}

func applyStyle(rows []string, style string) []string {
	switch style {
	case "normal":
		return applyNormal(rows)
	case "bold":
		return applyBold(rows)
	case "italic":
		return applyItalic(rows)
	case "outline":
		return applyOutline(rows)
	default:
		panic("unsupported style: " + style)
	}
}

func renderText(text string) []string {
	rows := make([]string, 8)
	for _, c := range text {
		pattern := basePattern(c)
		for i := 0; i < 8; i++ {
			rows[i] += pattern[i]
		}
	}
	return rows
}

func (ab *ArtBuilder) Build() string {
	if len(ab.segments) == 0 {
		return ""
	}

	finalRows := make([]string, 8)

	for _, seg := range ab.segments {
		rows := renderText(seg.text)
		rows = applyStyle(rows, seg.style)
		for i := 0; i < 8; i++ {
			finalRows[i] += rows[i]
		}
	}

	return strings.Join(finalRows, "\n") + "\n"
}