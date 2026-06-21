package main

import (
	"fmt"
	"time"
)

func serveTable(table int) {
	fmt.Printf("Starting table %d\n", table)
	time.Sleep(2 * time.Second)
	fmt.Printf("Finished table %d\n", table)
}

func main(){
	go serveTable(1)
	go serveTable(2)
	go serveTable(3)

	time.Sleep(3 * time.Second)
	fmt.Println("All done")
}