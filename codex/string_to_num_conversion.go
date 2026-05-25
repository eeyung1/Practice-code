package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]

	num := 0

	for _, r := range input {
		num = num*10 + int(r-'0')
	}

	fmt.Println(num)
}