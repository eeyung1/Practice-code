package main

import (
	"fmt"
	"os"
	"strconv"
)

func celsiusToFahrenheit(c float64) float64 {
    return (c * 9/5) + 32 
}

func fahrenheitToCelsius(f float64) float64 {
    return (f - 32) * 5/9
}

func main() {
    // step 1: confirm arguments (need exactly 2)
	if len(os.Args) != 3 {
		fmt.Println("Error: Invalid input")
		os.Exit(1)
	}

    // step 2: convert Os.Args[1] to a float64
	temp, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
    fmt.Println("Error: temperature must be a number")
    os.Exit(1)
	}

    // step 3: check Os.Args[2] for direction ("CtoF" or "FtoC")
    //         and call the right function

	var result float64  // declare here so everywhere can see it

	switch os.Args[2] {
	case "CtoF":
		result = celsiusToFahrenheit(temp)   // = not :=
	case "FtoC":
		result = fahrenheitToCelsius(temp)   // = not :=
	default:
		fmt.Println("Error: use CtoF or FtoC")
		os.Exit(1)
	}

    // step 4: output the result
	fmt.Printf("%.2f\n", result) 
}