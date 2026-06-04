package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) < 2 {
		return
	}

	args := os.Args[1:]

	for _, arg := range args {
		chars := []rune(arg)

		for i := 0; i < len(chars); i++ {
			isLastLetter := false

			if i == len(chars)-1 {
				isLastLetter = true
			} else if chars[i+1] == ' ' {
				isLastLetter = true
			}

			if isLastLetter {
				if chars[i] >= 'a' && chars[i] <= 'z' {
					chars[i] -= 32
				} else {
					if chars[i] >= 'A' && chars[i] <= 'Z' {
						chars[i] += 32
					}
				}
			}
		}

		fmt.Println(string(chars))
	}

}