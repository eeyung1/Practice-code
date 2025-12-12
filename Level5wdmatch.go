package main

import (
	"fmt"
	"os"
)

func wdMatch(str1, str2 string) {
	i, j := 0, 0

	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			i++
		}
		j++
	}

	if i == len(str1) {
		fmt.Println(str1)
	}
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	wdMatch(str1, str2)
}
