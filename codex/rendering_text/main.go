package main

import (
	"fmt"
)

func getCharLines(lines []string, index int) []string {
	start := index*4 + 1
	end := start + 3

	if start >= len(lines) || end > len(lines) {
		return []string{"", "", ""} 
	}

	return lines[start:end]
}

func renderWord(lines []string, word string) {
	for row := 0; row < 3; row++ {
		line := ""

		for _, char := range word {
			index := int(char) - 'A'
			getLines := getCharLines(lines, index)
			line += getLines[row]
		}
		fmt.Println(line)
	}
}

func main() {
	lines := []string{
		"",
		"A1", "A2", "A3",
		"",
		"B1", "B2", "B3",
		"",
		"C1", "C2", "C3",
	}

	renderWord(lines, "ABC")


	fmt.Println(getCharLines(lines, 0))
}