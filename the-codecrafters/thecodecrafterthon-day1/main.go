package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) float64 {
	return a / b
}

func pow(a, b float64) float64 {
	return math.Pow(a, b)
}

func mod(a, b int) int {
	return a%b
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		words := strings.Fields(input)

		if len(words) == 0 {
			continue
		}

		switch words[0] {
		case "quit":
			fmt.Println("goodbye")
			return

		case "help":
			fmt.Println("add <a> <b>   → addition")
			fmt.Println("sub <a> <b>   → subtraction")
			fmt.Println("mul <a> <b>   → multiplication")
			fmt.Println("div <a> <b>   → division")
			fmt.Println("pow <a> <b>   → power")
			fmt.Println("mod <a> <b>   → remainder of division")
			continue
		}

		if len(words) != 3 {
			fmt.Println("Invalid number of arguments")
			continue
		}

		a, err2 := strconv.ParseFloat(words[1], 64)
		if err2 != nil {
			fmt.Println("Invalid command type 'help'")
			continue
		}

		b, err1 := strconv.ParseFloat(words[2], 64)
		if err1 != nil {
			fmt.Println("Invalid number: must be numeric")
			continue
		}

		switch words[0] {
		case "add":
			fmt.Println(" ✦ Result:", add(a, b))
			continue
		case "sub":
			fmt.Println(" ✦ Result:", sub(a, b))
			continue
		case "mul":
			fmt.Println(" ✦ Result:", mul(a, b))
			continue
		case "div":
			if b == 0 {
				fmt.Println("Can't divide by zero")
				continue
			}
			fmt.Println(" ✦ Result:", div(a, b))
			continue
		case "pow":
			fmt.Println(" ✦ Result:", pow(a, b))
		case "mod":
			if b == 0 {
				fmt.Println("Invalid operation")
				continue
			}
			fmt.Println(" ✦ Result:", mod(int(a), int(b)))
		default:
			fmt.Println("Invalid command type 'help'")
			continue
		}

	}

}
