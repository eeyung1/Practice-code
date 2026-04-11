package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid number of arguments")
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	input := string(data)
	lines := strings.Split(input, "\n")

	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	nonEmpty := 0
	for _, line := range lines {
		if line != "" {
			nonEmpty++
		}
	}

	fmt.Println("Total lines:", len(lines))
	fmt.Println("Non-empty lines:", nonEmpty)
	fmt.Println()

	for i, line := range lines {
		fmt.Printf("%d: %s\n", i, line)
	}
}
