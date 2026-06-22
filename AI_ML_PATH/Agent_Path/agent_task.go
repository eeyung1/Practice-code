package main

import (
	"fmt"
	"sync"
	"time"
)

func fettchData(task string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Duration(len(task)) * 100 * time.Millisecond)

	ch <- "Result from " + task
}

func main(){
	tasks := []string{"Artist A", "Artist B", "Artist C"}

	ch := make(chan string)

	var wg sync.WaitGroup

	fmt.Println("Coordinator: Starting agents...")

	for _, task := range tasks {
		wg.Add(1)
		go fettchData(task, ch, &wg)
	}

	go func(){
		wg.Wait()
		close(ch)
		fmt.Println("Coordinator: All agents finished")
	}()

	fmt.Println("Coordinator: collecting results...")

	for result := range ch {
		fmt.Println("Coordinator received:", result)
	}

	fmt.Println("Coordinator: Done")
}