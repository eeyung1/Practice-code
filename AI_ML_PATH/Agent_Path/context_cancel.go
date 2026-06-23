package main

import (
	"context"
	"fmt"
	"time"
)

func Datafetch(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Agent %d: Cancelled! Stopping.\n", id)
			return
		default:
			fmt.Printf("Agent %d: Working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main(){
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	go Datafetch(ctx, 1)
	go Datafetch(ctx, 2)
	go Datafetch(ctx, 3)

	time.Sleep(3 * time.Second)
	fmt.Println("Main exiting")
}