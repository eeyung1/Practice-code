package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// step 1: confirm arguments
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run . 5 + 3")
		os.Exit(1)
	}

	// step 2: convert strings to numbers
	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("First argument must be a number")
		os.Exit(1)
	}

	operator := os.Args[2]

	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Third argument must be a number")
		os.Exit(1)
	}

	// step 3: logic
result, err := calculate(a, operator, b)
if err != nil {
    fmt.Println("Error:", err)
    os.Exit(1)
}

	// step 4: output
	fmt.Println(result)
}

func calculate(a float64, operator string, b float64) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b + 2, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", operator)
	}
}
