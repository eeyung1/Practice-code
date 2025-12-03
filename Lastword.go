package main

import (
	"fmt"
)

func LastWord(s string) string {
	for len(s) > 0 && s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}

	lastSpaceIndex := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			lastSpaceIndex = i
			break
		}
	}

	if lastSpaceIndex == -1 {
		return s
	}

	return s[lastSpaceIndex+1:]
}

func main() {
	fmt.Println(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Println(LastWord(" lorem,ipsum ")) // Note: The input has leading/trailing spaces
	fmt.Println(LastWord(" "))             // Returns ""
	fmt.Println(LastWord("SingleWord"))    // Test case for no spaces
}
