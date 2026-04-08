package main

import (
	"fmt"
)

func main(){
	lines := []string{
	"",
	"A1", "A2", "A3",
	"",
	"B1", "B2", "B3",
	"",
	"C1", "C2", "C3",
	}

	fmt.Println(getCharLines(lines, 0))
	fmt.Println(getCharLines(lines, 1))
	fmt.Println(getCharLines(lines, 2))
	fmt.Println(getCharLines(lines, 10))
}

func getCharLines(lines []string, index int) []string {
	start := index*4 + 1
	end := start + 3

	if end > len(lines) {
		return []string{}
	}

	return lines[start:end]
}