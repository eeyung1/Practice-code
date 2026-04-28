package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	inputText := os.Args[1]

	inputText = strings.ReplaceAll(inputText, "\\n", "\n")
	
	if inputText == "" {
		return
	}

	onlyNewlines := true

	for _, r := range inputText {
		if r != '\n' {
			onlyNewlines = false
			break
		}
	}

	if onlyNewlines {
		for i := 0; i < len(inputText); i++ {
			fmt.Println()
		}
		return
	}

	input := strings.Split(inputText, "\n")

	banner, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println("Error loading banner:", err)
		return
	}

	for i, chunk := range input {
		if chunk == "" {
			if i < len(input) - 1 {
				fmt.Println()
			}
			continue
		}

		lines := RenderLine(chunk, banner)

		for _, line := range lines {
			fmt.Println(line)
		}
	}

}
