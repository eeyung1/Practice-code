// func classify(input string) (string, error)
/*

It should return:
- `"empty"` if the string is `""`
- `"single character"` if length is 1
- `"short"` if length is 2–5
- `"medium"` if length is 6–15
- `"long"` if length is over 15
- An error if the string contains any character outside ASCII 32–126

In `main()`, test it with at least 6 different inputs including one that triggers the error.

*/

package main

import (
	"fmt"
	"errors"
)

func classify(input string) (string, error) {
	// ASCII validation FIRST
	for _, r := range input {
		if r < 32 || r > 126 {
			return "", errors.New("non-ASCII character detected")
		}
	}

	// Length classification
	length := len(input)

	if length == 0 {
		return "empty", nil
	} else if length == 1 {
		return "single character", nil
	} else if length >= 2 && length <= 5 {
		return "short", nil
	} else if length >= 6 && length <= 15 {
		return "medium", nil
	} else {
		return "long", nil
	}
}

func main() {
	tests := []string{
		"",
		"a",
		"word",
		"awesome",
		"rateofchemicalreaction",
		"hello\n", // triggers ASCII error
	}

	for _, t := range tests {
		result, err := classify(t)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(result)
		}
	}
}