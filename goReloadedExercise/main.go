package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var word string
	var operation string

	fmt.Println("----WELCOME TO STRING MANIPULATION TOOL----")
	time.Sleep(2 * time.Second)
	fmt.Println()

	fmt.Println("Input a word: ")
	fmt.Scanln(&word)

	for {
		fmt.Println("Select an operation to be carried out: (lastChar or capitalize or deleteIndex)")
		fmt.Scanln(&operation)

		if operation == "lastChar" || operation == "capitalize" || operation == "deleteIndex" {
			break
		}

		fmt.Println("Invalid operation. Please enter: lastChar, capitalize, or deleteIndex.")
	}

	if operation == "lastChar" {
		fmt.Println("The last character on the word is: ", word[len(word)-1:])
	}

	if operation == "capitalize" {
			fmt.Println("The capitalize word is: ", strings.ToUpper(word))
		}

		if operation == "deleteIndex" {
			var index int
			length := len(word)

			for {
				fmt.Println("Input an index: ")
				fmt.Scanln(&index)

				if index < 1 || index > length {
					fmt.Printf("Error: The index is from 0 to %d \n", length-1)
					continue
				}

				result := word[:index] + word[index+1:]
				fmt.Println("The Word without the character at the index inputted is: ", result)
				break
			}
		}
	
	
}
