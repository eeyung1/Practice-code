package main

import (
	"fmt"
)

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}

	if IsPrime(nb) {
		return nb
	}

	for i := nb; i >= 2; i-- {
		if IsPrime(i) {
			return i
		}
	}

	return 0
}

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}

	if nb == 2 {
		return true
	}

	if nb%2 == 0 {
		return false
	}

	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(FindPrevPrime(10))
}
