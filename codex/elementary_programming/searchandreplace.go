package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	text := os.Args[1]
	oldstr := os.Args[2]
	newstr := os.Args[3]

	result := strings.ReplaceAll(text, oldstr, newstr)

	fmt.Println(result)

	fmt.Println(ReplaceManual("aaaa", "aa", "x"))
	fmt.Println(CountReplace("banana", "na", "xo"))

}

func ReplaceManual(text, oldstr, newstr string) string {
	result := ""

	for i := 0; i < len(text); i++ {
		if i+len(oldstr) <= len(text) &&
			text[i:i+len(oldstr)] == oldstr {

			result += newstr
			i += len(oldstr) - 1
		} else {
			result += string(text[i])
		}
	}

	return result
}

func CountReplace(text, oldstr, newstr string) (string, int) {
		result := ""
		count := 0

	for i := 0; i < len(text); i++ {
		if i+len(oldstr) <= len(text) &&
			text[i:i+len(oldstr)] == oldstr {

			result += newstr
			i += len(oldstr) - 1
			count++
		} else {
			result += string(text[i])
		}
	}

	return result, count
}