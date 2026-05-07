package main

import "strings"

func reverseAscii(input string, patterns map[string]rune) string {
	lines := strings.Split(input, "\n")

	if len(lines) < 8 {
		return ""
	}

	result := ""

	width := len(lines[0])

	for col := 0; col < width; {
		found := false

		for size := 1; col+size <= width; size++ {

			blockLines := []string{}

			for row := 0; row < 8; row++ {
				blockLines = append(blockLines, lines[row][col:col+size])
			}

			block := strings.Join(blockLines, "\n")

			if ch, ok := patterns[block]; ok {
				result += string(ch)
				col += size
				found = true
				break
			}
		}

		if !found {
			col++
		}
	}

	return result
}