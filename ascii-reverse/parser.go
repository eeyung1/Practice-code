package main

import (
	"os"
	"strings"
)

func loadBanner(filename string) (map[string]rune, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	patterns := make(map[string]rune)

	ascii := 32

	for i := 1; i < len(lines); i += 9 {
		if i+8 > len(lines) {
			break
		}

		block := strings.Join(lines[i:i+8], "\n")

		patterns[block] = rune(ascii)

		ascii++
	}

	return patterns, nil
}