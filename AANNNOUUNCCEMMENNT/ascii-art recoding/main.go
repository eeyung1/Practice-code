package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) != 2 {
		return
	}

	filename := os.Args[1]

	fmt.Println(LoadBanner(filename))
}