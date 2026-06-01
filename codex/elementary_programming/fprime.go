package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Invalid number of arguments")
		return
	}

	arg := os.Args[1]

	n := 0

	for _, ch := range arg {

		if ch < '0' || ch > '9' {
			return
		}

		n = n*10 + int(ch-'0')

	}

	if n <= 1 {
		return
	}

	divisor := 2

	firstFactor := true

	for n > 1 {
		if n%divisor == 0 {
			if firstFactor {
				fmt.Print(divisor)
				firstFactor = false
			} else {
				fmt.Printf("*%d", divisor)
			}
			n /= divisor
		} else {
			divisor++
		}
	}

	fmt.Println()

}
