package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(" ════════════════════════════════════════════════")
	fmt.Println("SENTINEL — COMMAND & CONTROL CONSOLE")
	fmt.Println("All systems nominal. Type 'help' to begin.")
	fmt.Println(" ════════════════════════════════════════════════")

	start:
	fmt.Print("C&C> ")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	input := scanner.Text()

	cmd := strings.Fields(input)

	switch cmd[1] {
	case "upper":
		result := cmd
		result = result[2:]
		good := strings.Join(result, " ")
		fmt.Println("✦", strings.ToUpper(good))
		goto start

	case "title":
		result1 := cmd
		result1 = result1[2:]
		good1 := strings.Join(result1, " ")
		fmt.Println("✦", strings.Title(good1))
		goto start

	}

	if len(cmd) != 3 && len(cmd) != 4 {
		fmt.Println("Invalid number of arguments")
		goto start
	}

	if cmd[0] != "calc" && cmd[0] != "base" && cmd[0] != "str" {
		fmt.Println("Invalid command")
		goto start
	}

	if cmd[1] != "add" && cmd[1] != "sub" && cmd[1] != "mul" && cmd[1] != "div" && cmd[1] != "mod" && cmd[1] != "pow" && cmd[1] != "hex" && cmd[1] != "bin" && cmd[1] != "dec" {
		fmt.Println("Invalid command")
		goto start
	}
	switch cmd[1] {
	case "add":
		add1, err1 := strconv.ParseInt(cmd[2], 10, 64)
		if err1 != nil {
			fmt.Println("Invalid input")
			goto start
		}
		add2, err2 := strconv.ParseInt(cmd[3], 10, 64)
		if err2 != nil {
			fmt.Println("Invalid input")
			goto start 
		}

		fmt.Println("✦ Result:", add1+add2)
		goto start

	case "sub":
		sub1, err1 := strconv.ParseInt(cmd[2], 10, 64)
		if err1 != nil {
			fmt.Println("Invalid input")
			goto start
		}
		sub2, err2 := strconv.ParseInt(cmd[3], 10, 64)
		if err2 != nil {
			fmt.Println("Invalid Input")
			goto start
		}

		fmt.Println("✦ Result:", sub1-sub2)
		goto start
	case "mul":
		mul1, err1 := strconv.ParseInt(cmd[2], 10, 64)
		if err1 != nil {
			fmt.Println("Invalid Input")
			goto start
		}
		mul2, err2 := strconv.ParseInt(cmd[3], 10, 64)
		if err2 != nil {
			fmt.Println("Invalid Input")
			goto start
		}

		fmt.Println("✦ Result:", mul1*mul2)
		goto start
	case "div":
		div1, err1 := strconv.ParseInt(cmd[2], 10, 64)
		if err1 != nil {
			fmt.Println("Invalid Input")
			goto start
		}
		div2, err2 := strconv.ParseInt(cmd[3], 10, 64)
		if err2 != nil {
			fmt.Println("Invalid Input")
			goto start
		}

		if div2 == 0 {
			fmt.Println("Can't divide by zero")
			goto start 
		}

		fmt.Println("✦ Result:", float64(div1)/float64(div2))
		goto start
	case "mod":
		mod1, err1 := strconv.ParseInt(cmd[2], 10, 64)
		if err1 != nil {
			fmt.Println("Invalid Input")
			goto start
		}
		mod2, err2 := strconv.ParseInt(cmd[3], 10, 64)
		if err2 != nil {
			fmt.Println("Invalid Input")
			goto start
		}

		fmt.Println("✦ Result:", mod1%mod2)
		goto start
	case "pow":
		pow1, err1 := strconv.ParseInt(cmd[2], 10, 64)
		if err1 != nil {
			fmt.Println("Invalid Input")
			goto start
		}
		pow2, err2 := strconv.ParseInt(cmd[3], 10, 64)
		if err2 != nil {
			fmt.Println("Invalid Input")
			goto start
		}

		

		fmt.Println("✦ Result:", math.Pow(float64(pow1), float64(pow2)))
		goto start

	}

	switch cmd[1] {
	case "hex":
		hex1, err := strconv.ParseInt(cmd[2], 16, 64)
		if err != nil {
			fmt.Println("Invalid Input")
			goto start
		}

		fmt.Println("✦ Decimal: ", hex1)
		goto start

	case "bin":
		bin1, err := strconv.ParseInt(cmd[2], 2, 64)
		if err != nil {
			fmt.Println("Invalid Input")
			goto start
		}

		fmt.Println("✦ Decimal: ", bin1)
		goto start

	case "dec":
		dec1, err := strconv.ParseInt(cmd[2], 10, 64)
		if err != nil {
			fmt.Println("Invalid Input")
			goto start
		}

		fmt.Println("✦ Binary :", strconv.FormatInt(dec1, 2))
		fmt.Println("✦ Hex    :", strings.ToUpper(strconv.FormatInt(dec1, 16)))
		goto start
	}
	


}


