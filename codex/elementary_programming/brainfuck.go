package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	memory := make([]byte, 2048)
	ptr := 0
	code := os.Args[1]
	pc := 0

	for pc < len(code) {

		ch := code[pc]
		
		switch ch {
		case '>':
			ptr++
			if ptr >= 2048 {
				ptr = 0
			}
		case '<':
			ptr--
			if ptr < 0 {
				ptr = 2047
			}
		case '+':
			memory[ptr]++
		case '-':
			memory[ptr]--
		case '.':
			fmt.Printf("%c", memory[ptr])
		case '[':
			if memory[ptr] == 0 {
				count := 1
				for count > 0 {
					pc++ 
					if pc >= len(code) {
						fmt.Println("Error")
						return
					}

					if code[pc] == '[' {
						count++
					}

					if code[pc] == ']' {
						count--
					}
				}
			}

		case ']':
			if memory[ptr] != 0 {
				count := 1
				for count > 0 {
					pc--
					if pc < 0 {
						fmt.Println("Error")
						return
					}

					if code[pc] == ']' {
						count++
					}

					if code[pc] == '[' {
						count--
					}
				}
			}
		default:
			
		}

		pc++
	}
}