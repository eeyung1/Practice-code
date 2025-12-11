package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 1 {
		return
	}

	factors := []int{}
	divisor := 2

	for n > 1 {
		for n%divisor == 0 {
			factors = append(factors, divisor)
			n = n / divisor
		}
		divisor++
	}

	strFactors := []string{}
	for _, factor := range factors {
		strFactors = append(strFactors, strconv.Itoa(factor))
	}
	result := strings.Join(strFactors, "*")

	fmt.Println(result)
}
