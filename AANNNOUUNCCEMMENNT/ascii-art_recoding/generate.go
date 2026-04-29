package main

import (
	// "fmt"
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}

	segments := SplitInput(input)

	// Special Case: Check if input is ONLY literal \n sequences
	// Example: "\n\n" should result in two blank lines.
	onlyNewlines := true
	for _, seg := range segments {
		if seg != "" {
			onlyNewlines = false
			break
		}
	}

	var result strings.Builder

	if onlyNewlines {
		// If input was "\n", segments is ["", ""] (length 2)
		// We want to return one "\n" for each split except the last.
		for i := 0; i < len(segments)-1; i++ {
			result.WriteString("\n")
		}
		return result.String()
	}

	// Standard Case: Mixed text and newlines
	for i, seg := range segments {
		if seg == "" {
			// An empty segment inside the split represents a newline gap.
			// We skip the very last empty segment to avoid a trailing blank line.
			if i < len(segments)-1 {
				result.WriteString("\n")
			}
			continue
		}

		// Render the 8 rows for this segment
		rows := RenderLine(seg, banner)
		for _, row := range rows {
			result.WriteString(row + "\n")
		}
	}

	return result.String()
}
