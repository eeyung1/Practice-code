package main

import (
    "fmt"
)

func Slice(nbrs ...int) []int {
	return nbrs
}

func main(){
	fmt.Println(Slice(2, 3, 4, 5, 6))
}