package main

import "fmt"

func main() {
	err := ReadFile("example.txt")
	if err != nil {
		fmt.Println(err)
	}
}
