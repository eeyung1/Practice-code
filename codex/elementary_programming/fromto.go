package main

import (
	"fmt"
)

func FromTo(from int, to int) string {
	if from > 99 || to > 99 || from < 0 || to < 0 {
		return "Invalid\n"
	}

	result := ""
	if from <= to {
		for i := from; i <= to; i++ {

			result += fmt.Sprintf("%02d", i)

			if i != to {
				result += ", "
			}
		}
	} else {
		for i := from; i >= to; i-- {

			result += fmt.Sprintf("%02d", i)

			if i != to {
				result += ", "
			}
		}
	}

	return result + "\n"
}



func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(20, 5))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}