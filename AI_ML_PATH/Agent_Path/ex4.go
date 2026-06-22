package main

import (
	"fmt"
	"time"
)

func fetchArtist(name string, ch chan string) {
	time.Sleep(2 * time.Second)

	ch <- "Artist: " + name
}

func main(){
	ch := make(chan string)

	go fetchArtist("Burna Boy", ch)
	go fetchArtist("Wizkid", ch)

	result1 := <-ch
	result2 := <-ch

	fmt.Println(result1)
	fmt.Println(result2)
}