package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	str := os.Args[1]

	str = strings.TrimSpace(str)

	words := strings.Fields(str)

	if len(words) == 0 {
		return
	}

	result := strings.Join(words, "   ")

	fmt.Println(result)
}
