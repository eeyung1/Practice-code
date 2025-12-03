package main

import "fmt"

func FirstWord(s string) string {
	result := ""
	inWord := false

	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
			if inWord {
				break
			}
			continue
		}

		inWord = true
		result += string(r)
	}

	return result
}

func main() {
	fmt.Println(FirstWord("tfgugf hffyu there"))
	fmt.Println(FirstWord(""))
	fmt.Println(FirstWord("hello   .........  bye"))
	fmt.Println(FirstWord("   leading space"))
	fmt.Println(FirstWord("oneword"))
}
