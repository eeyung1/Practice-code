package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	args := os.Args[1:]

	
	for _, arg := range args {
		stack := []rune{}
		valid := true

		for _, ch := range arg {
			if ch == '(' || ch == '[' || ch == '{' {
				stack = append(stack, ch)
			} else if ch == ')' || ch == ']' || ch == '}' {
				if len(stack) == 0 {
					valid = false
					break
				}

				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if ch == ')' && top != '(' {
					valid = false
					break
				}

				if ch == ']' && top != '[' {
					valid = false
					break
				}

				if ch == '}' && top != '{' {
					valid = false
					break
				}
			}
		}

		if valid && len(stack) == 0 {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}

}