package main

import (
	"fmt"
	"os"
	"strings"
	//"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error: file doesn't exist")
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	characters := strings.ReplaceAll(string(data), "\r\n", "\n")

	lines := strings.Split(characters, "\n")

	banner := make(map[rune][]string)

	currentChar := rune(32)

	for i := 1; i < len(lines); i += 9 {
		if i+8 > len(lines) {
			break
		}

		charArt := lines[i : i+8]

		banner[currentChar] = charArt

		currentChar++
	}

	if len(banner) != 95 {
		return nil, fmt.Errorf("expected 95 characters, got %d", len(banner))
	}



	return banner, nil
}


// func LoadBanner(filename string) (map[rune][]string, error) {
// 	data, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}

// 	content := strings.ReplaceAll(string(data), "\r\n", "\n")
// 	lines := strings.Split(content, "\n")

// 	characters := make(map[rune][]string)
	
// 	currentChar := rune(32)
	
// 	for i := 1; i < len(lines); i += 9 {
// 		if i+8 > len(lines) {
// 			break
// 		}
		
// 		charLines := lines[i : i+8]
// 		characters[currentChar] = charLines
		
// 		currentChar++
// 	}
	
// 	return characters, nil
// }

// package main

// import (
// 	"os"
// 	"strings"
// )

// func LoadBanner(filename string) (map[rune][]string, error) {
// 	data, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, nil
// 	}

// 	lines := strings.Split(string(data), "\n")

// 	var characters map[rune][]string

// 	i := 1

// 	for i < len(lines) {
// 		charLines := make([]string, 8)

// 		for row := 0; row < 8; row++ {
// 			if i + row < len(lines) {
// 				charLines[row] = lines[i+row]
// 			} else {
// 				charLines[row] = ""
// 			}

// 			i += 9
// 		}

// 		for j := 32; j < 127; j++ {
// 			index := 32 + j

// 			characters[rune(index)] = charLines
// 		}
// 	}

// 	return characters, nil
// }



