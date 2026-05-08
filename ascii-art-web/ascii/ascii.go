package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GenerateArt takes the user's text and a banner name (e.g. "standard"),
// reads the corresponding banner file, and returns the ASCII art as a string.
func GenerateArt(text, banner string) (string, error) {
	// Build path to the banner file
	filePath := fmt.Sprintf("banners/%s.txt", banner)

	// Load all character definitions from the banner file
	chars, err := loadBanner(filePath)
	if err != nil {
		return "", err
	}

	var result strings.Builder

	// Each line of user input is processed separately
	// This handles multi-line input (user pressed Enter)
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		// An empty line in the input becomes an empty line in the output
		if line == "" {
			result.WriteString("\n")
			continue
		}

		// Each character in the banner is 8 rows tall.
		// We build the output row by row across all characters in the line.
		for row := 0; row < 8; row++ {
			for _, ch := range line {
				// ASCII printable characters start at space (32).
				// Each char occupies 8 lines in the banner file.
				// Formula: index of char's first line = (charCode - 32) * 9 + 1
				// The +1 skips the blank separator line before each character block.
				index := (int(ch)-32)*9 + 1 + row

				if index < 0 || index >= len(chars) {
					return "", fmt.Errorf("character '%c' (code %d) is not supported", ch, int(ch))
				}
				result.WriteString(chars[index])
			}
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}

// loadBanner reads a banner .txt file and returns all its lines as a slice.
// Each character in the banner is defined across 8 lines, separated by a blank line.
func loadBanner(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		// Distinguish between file-not-found and other errors
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("banner file not found: %s", filePath)
		}
		return nil, fmt.Errorf("could not open banner file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading banner file: %w", err)
	}

	return lines, nil
}
