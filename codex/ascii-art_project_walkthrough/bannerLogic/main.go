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

func main() {
	lines := []string{
		"",
		"A1", "A2", "A3",
		"",
		"B1", "B2", "B3",
	}

	fmt.Println(getCharBlock(lines, 0))
	fmt.Println(getCharBlock(lines, 1))
	fmt.Println(getCharBlock(lines, -1))
}
