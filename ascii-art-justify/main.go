package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 || len(args) > 3 {
		printUsage()
		return
	}

	alignType := "left"
	input := ""
	banner := "standard"

	argIndex := 0

	// Step 1: check for --align flag
	if strings.HasPrefix(args[0], "--align=") {
		parts := strings.SplitN(args[0], "=", 2)
		if len(parts) != 2 {
			printUsage()
			return
		}

		alignType = parts[1]

		// validate align type
		if alignType != "left" && alignType != "right" &&
			alignType != "center" && alignType != "justify" {
			printUsage()
			return
		}

		argIndex++
	}

	// Step 2: extract input
	if argIndex >= len(args) {
		printUsage()
		return
	}
	input = args[argIndex]
	argIndex++

	// Step 3: optional banner
	if argIndex < len(args) {
		banner = args[argIndex]
		argIndex++
	}

	// Step 4: ensure no extra arguments
	if argIndex != len(args) {
		printUsage()
		return
	}

	bannerMap, err := LoadBanner(banner)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ascii := Render(input, bannerMap)

	// Split into lines
	lines := strings.Split(ascii, "\n")

	// You can hardcode width for now
	width := 80

	var aligned []string

	switch alignType {
	case "left":
		aligned = AlignLeft(lines, width)
	case "right":
		aligned = AlignRight(lines, width)
	case "center":
		aligned = AlignCenter(lines, width)
	case "justify":
		aligned = AlignJustify(lines, width)
	}

	// Join back and print
	final := strings.Join(aligned, "\n")
	fmt.Print(final)
}

func printUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
}
