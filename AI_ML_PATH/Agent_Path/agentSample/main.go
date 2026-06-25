package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

func fetchData(ctx context.Context, id int, ch chan string, wg *sync.WaitGroup) {
    defer wg.Done()

    select {
    case <-ctx.Done():
        ch <- fmt.Sprintf("Agent %d: Cancelled", id)
        return
    case <-time.After(time.Duration(id) * time.Second):
        ch <- fmt.Sprintf("Agent %d: Completed", id)
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
    defer cancel()

    ch := make(chan string)
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go fetchData(ctx, i, ch, &wg)
    }

    go func() {
        wg.Wait()
        close(ch)
    }()

    for result := range ch {
        fmt.Println(result)
    }
}