package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func quotes(s string) string {
	lines := strings.Split(s, "\n")
	//fmt.Println(len(lines))
	var result []string

	for _, line := range lines {
		//fmt.Println(line)
		line = strings.TrimSpace(line)
		//fmt.Println(len(line))
		//fmt.Println(line)
		if line == "" {
			continue
		}

		// 1. Collapse multiple spaces into one
		elem := strings.Fields(line)
		//fmt.Println(len(elem))
		line = strings.Join(elem, " ")

		// 2. Fix punctuation: No space before, exactly one space after
		// This handles cases like "is ,the" -> "is, the"
		rePunct := regexp.MustCompile(`\s*([.,!?:;]+)\s*`)

		line = rePunct.ReplaceAllString(line, `$1 `)
		//fmt.Println(line)

		// 3. Fix quotes: Remove spaces inside double quotes
		reDouble := regexp.MustCompile(`"\s*(.*?)\s*"`)
		line = reDouble.ReplaceAllString(line, `"$1"`)

		// 4. Fix quotes: Remove spaces inside single quotes
		reSingle := regexp.MustCompile(`'\s*(.*?)\s*'`)
		line = reSingle.ReplaceAllString(line, `'$1'`)

		// 5. Final cleanup: Remove space before punctuation if any were added
		reClean := regexp.MustCompile(`\s+([,!?:;])`)
		line = reClean.ReplaceAllString(line, `$1`)
		//fmt.Println(len(line))

		result = append(result, strings.TrimSpace(line))
	}

	return strings.Join(result, "\n")
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	content := quotes(string(data)) + "\n"

	// Print the result to console as requested
	//fmt.Println(content)

	// Write the result to the output file
	err = os.WriteFile(os.Args[2], []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
	}
}
