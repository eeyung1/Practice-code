package ascii

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(bannerName string) ([][]string, error) {
	filePath := fmt.Sprintf("banners/%s.txt", bannerName)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read banner file '%s': %w", filePath, err)
	}

	lines := strings.Split(string(data), "\n")

	var characters [][]string

	i := 1
	for i < len(lines) {
		charLines := make([]string, 8)
		for row := 0; row < 8; row++ {
			if i+row < len(lines) {
				charLines[row] = lines[i+row]
			} else {
				charLines[row] = ""
			}
		}
		characters = append(characters, charLines)
		i += 9
	}

	return characters, nil
}

func Render(input string, bannerName string) (string, error) {
	if input == "" {
		return "", nil
	}
	if input == "\n" {
		return "\n", nil
	}
	banner, err := LoadBanner(bannerName)
	if err != nil {
		return "", err
	}

	lines := strings.Split(input, "\n")
	var result strings.Builder

	for _, line := range lines {
		if line == "" {
			result.WriteByte('\n')
			continue
		}
		for row := range 8 {
			for _, ch := range line {
				index := int(ch) - 32
				if index < 0 || index >= len(banner) {
					continue
				}

				result.WriteString(banner[index][row])
			}
			result.WriteByte('\n')
		}
	}
	return result.String(), nil
}
