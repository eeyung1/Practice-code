package main

import (
	"fmt"
)

func getCharBlock(lines []string, index int) []string {

	if index < 0 {
		return []string{}
	}

	blockSize := 4
	charHeight := 3

	start := index*blockSize + 1
	end := start + charHeight

	if start >= len(lines) || end > len(lines) {
		return []string{}
	}

	return lines[start:end]
}

func renderWord(lines []string, word string) {
	charHeight := 3

	for i := 0; i < charHeight; i++ {
		line := ""

		for _, r := range word {
			index := int(r) - 'A'
			charLines := getCharBlock(lines, index)

			if len(charLines) == charHeight {
				line += charLines[i]
			} else {
				line += " "
			}
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

	renderWord(lines, "AB")
	//renderWord(lines, "ABC")

	fmt.Println(getCharBlock(lines, 0))
	fmt.Println(getCharBlock(lines, 1))

}
