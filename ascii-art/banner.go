package main

import (
	"fmt"
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	removeNewline := strings.ReplaceAll(string(fileBytes), "\r\n", "\n")
	lines := strings.Split(removeNewline, "\n")

	currentCharacter := rune(32)

	banner := make(map[rune][]string)

	for index := 1; index < len(lines); index += 9 {
		if index+8 > len(lines) {
			continue
		}

		charArt := lines[index : index+8]

		banner[currentCharacter] = charArt

		currentCharacter++
	}

	if len(banner) != 95 {
		return nil, fmt.Errorf("expected 95 characters but got %d", len(banner))
	}

	return banner, nil
}
