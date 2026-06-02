package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]

	s2 := os.Args[2]

	position := 0

	for _, ch := range s2 {
		if position < len(s1) && ch == rune(s1[position]) {
			position++
		}
	}

	if position == len(s1) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}

}
