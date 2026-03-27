package main

import (
	"fmt"
)

func main(){
	for {
		fmt.Println("===CLI=CALCULATOR===")
		fmt.Println()

		start:
		fmt.Println("Select Operations by inputting the number")
		fmt.Println(" [1] Addition\n [2] Subtraction\n [3] Multiplication\n [4] Division\n [5] help\n [6] quit")

		var op int
		fmt.Scanln(&op)



		if op == 1 {
			var a float64
			var b float64

			var operation string

			fmt.Println("add <a> <b>")
			_, err := fmt.Scanln(&operation, &a, &b)

			if operation != "add" {
				fmt.Println("Invalid Operation select help for more information")
				goto start
			} 
			
			if err != nil {
				fmt.Println("Invalid Operation select help for more information")
				goto start
			}

			fmt.Println("Result: ", a + b)
			fmt.Println()
			goto start
		}

		if op == 2 {
			var a float64
			var b float64

			var operation string

			fmt.Println("sub <a> <b>")
			_, err := fmt.Scanln(&operation, &a, &b)

			if operation != "sub" {
				fmt.Println("Invalid Operation select help for more information")
				goto start
			} 
			
			if err != nil {
				fmt.Println("Invalid Operation select help for more information")
				goto start
			}

			fmt.Println("Result: ", a - b)
			fmt.Println()
			goto start
		}


		if op == 3 {
			var a float64
			var b float64

			var operation string

			fmt.Println("mul <a> <b>")
			_, err := fmt.Scanln(&operation, &a, &b)

			if operation != "mul" {
				fmt.Println("Invalid Operation select help for more information")
				goto start
			} 
			
			if err != nil {
				fmt.Println("Invalid Operation select help for more information")
				goto start
			}

			fmt.Println("Result: ", a * b)
			fmt.Println()
			goto start
		}

		if op == 4 {
			var a float64
			var b float64

			var operation string

			fmt.Println("div <a> <b>")
			_, err := fmt.Scanln(&operation, &a, &b)

			if operation != "div" {
				fmt.Println("Invalid Operation select help for more information")
				goto start
			} 
			
			if err != nil {
				fmt.Println("Invalid Operation select help for more information")
				goto start
			}

			if b == 0 {
				fmt.Println("Can't divide by zero")
				goto start
			}

			fmt.Println("Result: ", a / b)
			fmt.Println()
			goto start
		}

		if op == 5 {
			fmt.Println("Here's the guidelines on how to use this calculator")
			fmt.Println()
			fmt.Println("(1) To add, Input 1 and type add followed by the numbers e.g add 4 5")
			fmt.Println()
			fmt.Println("(2) To subtract, Input 2 and type sub followed by the numbers e.g sub 6 5")
			fmt.Println()
			fmt.Println("(3) To multiple, Input 3 and type mul followed by the numbers e.g mul 4 5")
			fmt.Println()
			fmt.Println("(4) To divide, Input 4 and type div followed by the numbers e.g div 4 5")
			fmt.Println()
		}

		if op == 6 {
			fmt.Println("Goodbye and thanks for your time")
			fmt.Println()
			break
		}

		if op != 1 && op != 2 && op != 3 && op != 4 && op != 5 && op != 6 {
			fmt.Println("Select 5 for help guidelines")
			fmt.Println()
			goto start
		}
	


		continue
	}
}