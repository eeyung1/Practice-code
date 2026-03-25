package main

import (
	"fmt"
)

func Multiply(n int) string {
	result := ""
	for i := 3; i <= 24; i++ {
		result += fmt.Sprintf("%d x %d = %d\n", n, i, (i*n))
	}

	return result
}

// func main(){
// 	fmt.Println(Multiply(35))
// }