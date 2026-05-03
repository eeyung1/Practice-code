package main

import (
	"os"
	"fmt"
	"strings"
)

func LoadBanner(name string) (map[rune][]string, error) {
	filepath := fmt.Sprintf("%s.txt", name)
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("could not read banner file %s: %w", filepath, err)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("empty file read")
	}

	characters := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(characters, "\n")

	banner := make(map[rune][]string)

	firstChar := rune(32)

	for i := 1; i < len(lines); i += 9 {
		if i + 7 >= len(lines) {
			break
		}

		charArt := lines[i : i+8]
		banner[firstChar] = charArt
		firstChar++
	}

	if len(banner) != 95 {
		return nil, fmt.Errorf("expected 95 but got %d", len(banner))
	}

	return banner, nil
}