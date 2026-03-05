package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main(){
	if len(os.Args) != 2 {
		fmt.Println("Error reading input through the CLI")
		os.Exit(1)
	}

	inputFile := os.Args[1]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Problem reading the inputFile")
		os.Exit(1)
	}

	text := string(data)

	fmt.Println("File:", inputFile)
	fmt.Println("Size:", countChars(text), "characters")
	fmt.Println("Lines:", countLines(text))
	fmt.Println("Words:", countWords(text))
	fmt.Println("Extension:", filepath.Ext(inputFile))
}

func countChars(text string) int {
	return len(text)
}

func countLines(text string) int {
    text = strings.TrimRight(text, "\n")
    lines := strings.Split(text, "\n")  // ← missing this line!
    return len(lines)
}

func countWords(text string) int {
	    words := strings.Fields(text)
    return len(words)
}