package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [OPTION]")
		fmt.Println()
		fmt.Println("EX: go run . --reverse=<fileName>")
		return
	}

	arg := os.Args[1]

	if !isReverseFlag(arg) {
		fmt.Println("Usage: go run . [OPTION]")
		fmt.Println()
		fmt.Println("EX: go run . --reverse=<fileName>")
		return
	}

	fileName := getFileName(arg)

	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	patterns, err := loadBanner("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	result := reverseAscii(string(data), patterns)

	fmt.Println(result)
}