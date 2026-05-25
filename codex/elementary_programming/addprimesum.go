package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}

	num := os.Args[1]

	if num[0] == '-' {
		fmt.Println(0)
		return
	}

	n := 0
	sum := 0

	for _, r := range num {
		if r < '0' || r > '9' {
			fmt.Println(0)
			return
		} else {
		n = n*10 + int(r-'0')
		}
	}


	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}


	fmt.Println(sum)

}

func isPrime(n int)bool{
	if n < 2 {
		return false
	}

	if n == 2 {
		return true
	}

	if n % 2 == 0 {
		return false
	}

	for i := 3; i * i <= n; i += 2 {
		if n % i == 0 {
			return false
		}
	}

	return true
}
