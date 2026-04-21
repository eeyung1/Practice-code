package main

import (
	"fmt"
)

func main(){
	var colourMap = map[string]string{
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	}

	const reset = "\033[0m"

	fmt.Println(colourMap["red"] + "Hello World" + colourMap["green"] + " from golang" + reset)
	fmt.Println(colourMap["cyan"] + "PHYSICS" + reset)

}