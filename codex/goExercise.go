package main

import (
	"fmt"
)

func main(){
	var Name string = "Alice"
	var Age int = 25

	Height := 1.75
	Is_student := true

	fmt.Printf("Name: %s\n", Name)
	fmt.Printf("Age: %d\n", Age)
	fmt.Printf("Height: %2f\n", Height)
	fmt.Printf("Is student: %t\n", Is_student)
}