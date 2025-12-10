package main

import (
	"fmt"
	"strings"
	"unicode" // Import the unicode package for character checks
)

func main() {
	var name string
	fmt.Println("Enter your name:")

	// 1. Check for basic input errors from Scanln
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return // Exit the program if input reading fails
	}

	name = strings.TrimSpace(name) // Clean up input

	// 2. Validate the *content* of the input (ensure it's not empty)
	if name == "" {
		fmt.Println("Error: Name cannot be empty.")
		return
	}

	// 3. Check for the presence of numbers
	for _, character := range name {
		// unicode.IsDigit(r) returns true if the character is a number (0-9)
		if unicode.IsDigit(character) {
			fmt.Println("Error: Name cannot contain numbers. Please use letters only.")
			return // Exit immediately because the input is invalid
		}
	}

	// If we reach this point, the name is valid
	fmt.Println("Hello! ", name)
}
