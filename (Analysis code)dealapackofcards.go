package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		fmt.Printf("Player %d: %d, %d, %d\n", i+1, deck[i*3], deck[i*3+1], deck[i*3+2])
	}
}
