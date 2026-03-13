package main

import (
	"fmt"
)

func Multiply(n int) string {
	result := ""
	for i := 1; i <= 70; i++ {
		result += fmt.Sprintf("%d x %d = %d\n", n, i, (i*n))
	}

	return result
}

func main(){
	fmt.Println(Multiply(12))
}