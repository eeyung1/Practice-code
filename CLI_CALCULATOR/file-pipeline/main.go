package main

import (
	"fmt"
	"strings"
	"os"
)

func CapsToTitle(word string) string {
	return strings.Title(word)
}

func CLASSIFIED(word string) string {
	return strings.ReplaceAll(word, "CLASSIFIED", "[REDACTED]")
}

func lowerToUpper(word string) string {
	return strings.ToUpper(word)
}

func reverse(word string) string {
	if len(word) == 0 {
		return ""
	}

	good := strings.Fields(word)

	for i := 0; i < len(good); i++ {
		good[i] = reverse(good[i][1:]) + string(good[i][0])
	}

	return strings.Join(good, " ")
}

func replaceTodo(word string) string {
	word = strings.ReplaceAll(word, "TODO:", "✦ ACTION:")
	return word
}

func main(){
	if len(os.Args) != 3 {
		fmt.Println("")
	}
}

