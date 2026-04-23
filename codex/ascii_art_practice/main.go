package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <string>")
		return
	}

	input := os.Args[1]
	if input == "" {
		return
	}

	fmt.Print(AsciiArt(input))
}

func AsciiArt(input string) string {
	banner, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	bannerStr := strings.ReplaceAll(string(banner), "\r\n", "\n")
	bannerLines := strings.Split(bannerStr, "\n")

	words := strings.Split(input, "\\n")

	var result strings.Builder

	for _, word := range words {
		if word == "" {
			result.WriteByte('\n')
			continue
		}

		for i := 0; i < 8; i++ {
			for _, char := range word {
				startLine := (int(char-' ') * 9) + 1
				result.WriteString(bannerLines[startLine+i])
			}
			result.WriteByte('\n')
		}
	}

	return result.String()
}
