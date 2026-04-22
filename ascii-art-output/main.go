package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		fmt.Print(Render("standard.txt", args[0]))
		return
	}

	if len(args) != 3 || !strings.HasPrefix(args[0], "--output=") {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("\nEX: go run . --output=<fileName.txt> something standard")
		return
	}

	outputArg := args[0]
	parts := strings.SplitN(outputArg, "=", 2)
	if len(parts) < 2 || parts[1] == "" {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}
	outputFile := parts[1]

	bannerFile := args[2] + ".txt"
	input := args[1]

	result := Render(bannerFile, input)

	err := os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
	}
}

func Render(bannerPath, input string) string {
	data, err := os.ReadFile(bannerPath)
	if err != nil {
		return ""
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	bannerLines := strings.Split(content, "\n")

	input = strings.ReplaceAll(input, "\\n", "\n")
	words := strings.Split(input, "\n")

	var result strings.Builder

	for _, word := range words {
		if word == "" {
			result.WriteString("\n")
			continue
		}

		for i := 1; i <= 8; i++ {
			for j := 0; j < len(word); j++ {
				index := int(word[j]-32)*9 + i
				if index < len(bannerLines) {
					result.WriteString(bannerLines[index])
				}
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}