package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	printed := map[rune]bool{}
	
	for _, ch := range s1 {
		if printed[ch] {
			continue
		}

		found := false

		for _, c := range s2 {
			if ch == c {
				found = true
				break
			}
		}

		if found {
			fmt.Print(string(ch))
			printed[ch] = true
		}
	}

	fmt.Println()
}