package main

import (
	"fmt"
	"sync"
)

func feetchData(task string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	ch <- "Result from " + task
}

func main(){
	tasks := []string{"Artist A", "Artist B", "Artist C", "Artist D"}

	ch := make(chan string)

	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)

		go feetchData(task, ch, &wg)
	}

	go func(){
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}
}