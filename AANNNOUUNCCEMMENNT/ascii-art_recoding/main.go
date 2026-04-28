package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	inputText := os.Args[1]

	banner, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println("Error loading banner:", err)
		return
	}

	lines := Renderlines(inputText, banner)

	for _, line := range lines {
		fmt.Println(line)
	}
}
