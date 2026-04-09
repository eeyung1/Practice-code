package main

import (
	"fmt"
	"os"
	"strings"
)

func getCharLines(lines []string, char rune) []string {
	index := int(char) - 32
	start := index * 9
	end := start + 8

	if start >= len(lines) || end > len(lines) {
		return []string{"", "", "", "", "", "", "", ""}
	}

	return lines[start:end]
}

func renderWord(lines []string, word string) {
	for row := 0; row < 8; row++ {
		line := ""

		for _, char := range word {
			charLines := getCharLines(lines, char)
			line += charLines[row]
		}

		fmt.Println(line)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <text>")
		return
	}

	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	clean := strings.ReplaceAll(string(data), "\r", "")
	lines := strings.Split(clean, "\n")

	input := os.Args[1]
	words := strings.Split(input, "\\n")

	for _, word := range words {
		renderWord(lines, word)
	}
}