package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchData1(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(id) * time.Second)

	fmt.Printf("Agent %d finished\n", id)
}

func main(){
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go fetchData1(i, &wg)
	}

	wg.Wait()
	fmt.Println("All agents finished")
}