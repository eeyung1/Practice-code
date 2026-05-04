package main

import "strings"

// LEFT → no change
func AlignLeft(lines []string, width int) []string {
	return lines
}

// RIGHT → pad left
func AlignRight(lines []string, width int) []string {
	result := make([]string, len(lines))

	for i, line := range lines {
		if len(line) >= width {
			result[i] = line
			continue
		}

		padding := width - len(line)
		result[i] = strings.Repeat(" ", padding) + line
	}

	return result
}

// CENTER → split padding left/right
func AlignCenter(lines []string, width int) []string {
	result := make([]string, len(lines))

	for i, line := range lines {
		if len(line) >= width {
			result[i] = line
			continue
		}

		padding := width - len(line)
		left := padding / 2
		right := padding - left

		result[i] = strings.Repeat(" ", left) + line + strings.Repeat(" ", right)
	}

	return result
}

// JUSTIFY → distribute spaces between words
func AlignJustify(lines []string, width int) []string {
	result := make([]string, len(lines))

	for i, line := range lines {

		// If line too long or no spaces → keep as is
		if len(line) >= width || !strings.Contains(line, " ") {
			result[i] = line
			continue
		}

		words := strings.Fields(line)

		// If only one word → same as left
		if len(words) == 1 {
			result[i] = line
			continue
		}

		// Total length of words (no spaces)
		totalChars := 0
		for _, w := range words {
			totalChars += len(w)
		}

		// Total spaces needed
		totalSpaces := width - totalChars

		// Number of gaps between words
		gaps := len(words) - 1

		// Minimum spaces per gap
		spacePerGap := totalSpaces / gaps

		// Extra spaces to distribute
		extra := totalSpaces % gaps

		var builder strings.Builder

		for j, word := range words {
			builder.WriteString(word)

			// Add spaces after each word except last
			if j < gaps {
				spaces := spacePerGap

				// distribute extra spaces from left to right
				if j < extra {
					spaces++
				}

				builder.WriteString(strings.Repeat(" ", spaces))
			}
		}

		result[i] = builder.String()
	}

	return result
}