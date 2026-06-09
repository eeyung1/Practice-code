package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	position := 0
	result := ""

	for _, ch := range str1 {
		found := false

		for i := position; i < len(str2); i++ {
			if str2[i] == byte(ch) {
				found = true
				position = i + 1
				break
			}
		}

		if !found {
			return
		}

		result += string(ch)
	}

	fmt.Println(result)
}