package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}

	exp := os.Args[1]

	tokens := strings.Fields(exp)

	stack := []int{}

	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack = append(stack, num)
		} else {
			if token == "+" || token == "-" || token == "*" || token == "/" || token == "%" {
				if len(stack) < 2 {
					fmt.Println("Error")
					return
				}

				b := stack[len(stack)-1]
				a := stack[len(stack)-2]


				stack = stack[:len(stack)-2]

				var result int

				switch token {
				case "+":
					result = a + b
				case "-":
					result = a - b
				case "*":
					result = a * b
				case "/":
					result = a / b
				case "%":
					result = a % b
				}

				stack = append(stack, result)
			} else {
				fmt.Println("Error")
				return
			}
		}
	}

	if len(stack) != 1 {
		fmt.Println("Error")
		return
	}

	fmt.Println(stack[0])
}

