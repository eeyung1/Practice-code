package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)

	start:

	fmt.Print("> ")

	input, _ := reader.ReadString('\n')

	word := strings.Fields(input)

	if word[0] == "quit" {
		fmt.Println("Goodbye")
		return
	}

	if len(word) != 3 {
		fmt.Println("Invalid number of arguments")
		goto start
	}

	if word[0] != "convert" {
		fmt.Println("Invalid command")
		goto start
	}

	if word[2] != "hex" && word[2] != "bin" && word[2] != "dec" {
		fmt.Println("Invalid command")
		goto start
	} 

	switch word[2] {
	case "hex":
		great, _ := strconv.ParseInt(word[1], 16, 64)

		fmt.Println("✦ Decimal: ", strconv.FormatInt(great, 10))
		goto start

	case "bin":
		ok, _ := strconv.ParseInt(word[1], 2, 64)

		fmt.Println("✦ Decimal: ", strconv.FormatInt(ok, 10))
		goto start

	case "dec":
		ok1, _ := strconv.ParseInt(word[1], 10, 64)

		fmt.Println("✦ Binary: ", strconv.FormatInt(ok1, 2))
		fmt.Println("✦ Hex: ", strings.ToUpper(strconv.FormatInt(ok1, 16)))
		goto start
	}

}