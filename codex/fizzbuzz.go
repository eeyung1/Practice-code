package main

import (
	"fmt"
)

func FizzBuzz(n int) string {
	result := ""
	for i := 1; i <= n; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			result += "FizzBuzz\n"
		} else if i % 3 == 0  {
			result += "Fizz\n"
		} else if i % 5 == 0 {
			result += "Buzz\n"
		} else {
			result += fmt.Sprintf("%d\n", i)
		}
	}

	return result
}

// func main (){
// 	fmt.Println(FizzBuzz(15))
// }