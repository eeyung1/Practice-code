package main

import (
	"fmt"
)

func main(){
	A := 'A'
	fmt.Printf("A as a number: %v\n", A)
	fmt.Printf("A as a character: %v\n", string(A))
	fmt.Printf("Next character: %v\n", string(A + 1))

	value := 122
	fmt.Printf("122 as character: %v\n", string(value))

	word := "Hello"
	for _, r := range word {
		fmt.Printf("%c = %v\n", r, r)
	}
}