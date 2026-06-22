package main

import (
	"fmt"
	"time"
)

func fetchData(agentName string, ch chan string) {
	time.Sleep(2 * time.Second)

	ch <- agentName + " returned data"
}

func main(){
	ch := make(chan string)

	go fetchData("Agent A", ch)
	go fetchData("Agent B", ch)
	go fetchData("Agent C", ch)


	result1 := <-ch
	result2 := <-ch
	result3 := <-ch

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	time.Sleep(3 * time.Second)
}