package main

import (
	"fmt"
)

// func FirstWord(s string) string {
// 	if s == "" {
// 		return "\n"
// 	}

// 	word := strings.Fields(s)

// 	return word[0] + "\n"
// }

func isSpace(c byte) bool {
	return c == ' ' || c == '\t'
}

func FirstWord(s string) string {
	i := 0

	//skip leading spaces
	for i < len(s) && isSpace(s[i]) {
		i++
	}

	start := i

	//collect word
	for i < len(s) && !isSpace(s[i]) {
		i++
	}

	if start == i {
		return "\n"
	}

	return s[start:i] + "\n"
}

func LastWord(s string) string {
	i := len(s) - 1

	// skip trailing spaces
	for i >= 0 && isSpace(s[i]) {
		i--
	}

	if i < 0 {
		return "\n"
	}

	end := i

	// move to start of word
	for i >= 0 && !isSpace(s[i]) {
		i--
	}

	return s[i+1 : end+1] + "\n"
}

func FirstAndLastWord(s string) string {
	a := FirstWord(s)
	b := LastWord(s)

	return a[:len(a)-1] + " " + b
	
}



func main() {
	fmt.Print(FirstAndLastWord("   hello world  test " ))
	fmt.Print(FirstAndLastWord("hello" ))
	fmt.Print(LastWord("hello world"))
	fmt.Print(LastWord("   hello   world  "))
	fmt.Print(LastWord(""))
	fmt.Print(LastWord("   "))
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
	fmt.Print(FirstWord("   hello\tworld"))
}
