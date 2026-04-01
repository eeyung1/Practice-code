package main

import (
	"bufio"
	"fmt"
	"os"
)

var lastResult string

var hasLast bool
var History []string

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a name: ")

		if !scanner.Scan() {
			break
		}



		input := scanner.Text()
		
		if input == "exit" {
			break
		}

		History = append(History, input)
		continue
	}

	for i := 0; i < len(History); i++ {
		fmt.Println(History[i])
	}

}
