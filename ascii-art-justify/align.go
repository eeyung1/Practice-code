package main

import "strings"

// Helper: get the longest line width
func getMaxWidth(lines []string) int {
	max := 0
	for _, line := range lines {
		if len(line) > max {
			max = len(line)
		}
	}
	return max
}

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

// JUSTIFY → fallback (real justify must happen BEFORE rendering)
func AlignJustify(lines []string, width int) []string {
	// ASCII-art rows cannot be safely justified here
	// because spacing must be applied between words BEFORE rendering

	return lines
}