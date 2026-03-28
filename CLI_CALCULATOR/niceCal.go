package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]

		switch cmd {
		case "help":
			printHelp()
			continue
		case "quit":
			fmt.Println("Goodbye")
			return
		}

		if len(parts) != 3 {
			fmt.Println("Invalid input. Type 'help'")
			continue
		}

		a, err1 := strconv.ParseFloat(parts[1], 64)
		b, err2 := strconv.ParseFloat(parts[2], 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Invalid numbers. Type 'help'")
			continue
		}

		result, err := calculate(cmd, a, b)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("✦ Result:", result)
	}
}

func calculate(cmd string, a, b float64) (float64, error) {
	switch cmd {
	case "add":
		return a + b, nil
	case "sub":
		return a - b, nil
	case "mul":
		return a * b, nil
	case "div":
		if b == 0 {
			return 0, fmt.Errorf("Can't divide by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("Unknown command. Type 'help'")
	}
}

func printHelp() {
	fmt.Println("Guidelines:")
	fmt.Println(" add <a> <b>  → addition")
	fmt.Println(" sub <a> <b>  → subtraction")
	fmt.Println(" mul <a> <b>  → multiplication")
	fmt.Println(" div <a> <b>  → division")
	fmt.Println(" help         → show this guide")
	fmt.Println(" quit         → exit program")
}