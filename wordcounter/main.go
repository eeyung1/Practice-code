package main

import (
    "fmt"
    "os"
    "strings"
)

func countWords(text string) int {
    words := strings.Fields(text)
    return len(words)
}

func countLines(text string) int {
    text = strings.TrimRight(text, "\n")  // remove trailing newline
    lines := strings.Split(text, "\n")
    return len(lines)
}

func countChars(text string) int {
    text = strings.TrimRight(text, "\n")  // remove trailing newline
    return len(text)
}

func main() {
    // step 1: confirm arguments
	if len(os.Args) != 2 {
		fmt.Println("Error")
		os.Exit(1)
	}
    // step 2: read the file
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
    fmt.Println("Error reading file:", err)
    os.Exit(1)
}
    // step 3: call the three functions

	text := string(data)

	lines := countLines(text)
	words := countWords(text)
	chars := countChars(text)
    
	// step 4: print the results

	fmt.Println("Lines:", lines) 
	fmt.Println("Words:", words) 
	fmt.Println("Characters:", chars)

}

