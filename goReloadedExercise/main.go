package main

import (
	"fmt"
	"strings"
)

func main(){
	var word string
	var operation string

	fmt.Println("------WELCOME TO STRING MANIPULATION TOOL------")

	fmt.Println("Input a word:")
	fmt.Scanln(&word)

	fmt.Println("Input an operation to be carried out: either lastletter, capitalize or deleteIndex: ")
	fmt.Scanln(&operation)

	if operation == "lastletter" {
		fmt.Println("The last letter of the word is: ", word[len(word)-1:])
	} else if operation == "capitalize" {
		fmt.Println("The capitalize word is: ", strings.ToUpper(word) )
	}

	if operation == "deleteIndex" {
		var index int
		length := len(word)

		for {
			fmt.Println("Input an index: ")
			fmt.Scanln(&index)

			if index < 1 || index > length {
				fmt.Printf("The word range from 0 to %d \n", length-1)
				continue
			}

			result := word[:index] + word[index + 1:]
			fmt.Println(result)
			break
		}
	}

}