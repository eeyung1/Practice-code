package main

import (
	"fmt"
)

func main() {
	inWord := false

	for _, ch := range "yes no ok" {
		if ch == ' ' {
			if !inWord {
				fmt.Println("start of word")
			}

			inWord = true
		} else {
			inWord = false
		}
	}
}

// func main() {
// 	for _, ch := range "a b" {
// 		if ch != ' ' {
// 			fmt.Println("letter")
// 		}
// 	}
// }
