package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// LoadBanner - Keep this here so the test can find it
func LoadBanner(filename string) (map[rune][]string, error) {
	val, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	data := strings.ReplaceAll(string(val), "\r\n", "\n")
	lines := strings.Split(data, "\n")
	storemap := make(map[rune][]string)
	asciiRune := rune(32)
	for i := 1; i+8 < len(lines); i += 9 {
		storemap[asciiRune] = lines[i : i+8]
		asciiRune++
	}
	if len(storemap) == 0 {
		return nil, errors.New("invalid format")
	}
	return storemap, nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		return
	}

	inputText := os.Args[1]
	bannerFile := os.Args[2]

	bannerMap, err := LoadBanner(bannerFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	inputLines := strings.Split(inputText, "\\n")
	for _, line := range inputLines {
		if line == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range line {
				if art, exists := bannerMap[char]; exists {
					fmt.Print(art[i])
				}
			}
			fmt.Println()
		}
	}
}
