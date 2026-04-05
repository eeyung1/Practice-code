// package main

// import "fmt"

// "fmt"
// "strings"

// func main(){
// 	/*
// 	fruits := []string{"apple", "banana", "cherry"}

// 	// Adding to a slice with append
// 	fruits = append(fruits, "date", "cashew")
// 	fmt.Println(fruits) // [apple banana date cherry]

// 	fmt.Println(len(fruits))
// 	fmt.Println(fruits[0])   // "apple"
// 	fmt.Println(fruits[1])

// 	fmt.Println(fruits[1:3]) // [banana cherry]  (index 1 up to but not including 3)
// 	fmt.Println(fruits[:2])  // [apple banana]   (from start to index 2)
// 	fmt.Println(fruits[2:])

// 	*/

// 	/*

// 	var sb strings.Builder

// 	sb.WriteString("Hello")
// 	sb.WriteString(", ")
// 	sb.WriteString("World")

// 	result := sb.String() // "Hello, World"
// 	fmt.Println(result)

// 	*/

// }


// package main

// import (
// 	"fmt"
// )

// func main(){
// 	sum := 1
// 	for ; sum < 1000; {
// 		sum += sum
// 	}

// 	fmt.Println(sum)
// }


package main

import (
	"fmt"
)

func main(){
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}