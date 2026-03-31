package main

import (
	"fmt"
	"os"
	"bufio"
)

var lastResult string

var hasLast bool
var History []string

func main(){
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a name: ")

	scanner.Scan()

	input := scanner.Text()

	fmt.Println(input)
}