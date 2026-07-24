package main

import "fmt"

type Room struct {
	Name 	string
	X    	int
	Y    	int
	IsStart	bool
	IsEnd	bool
}

func (r Room) Display() {
	fmt.Printf("Room: %s\n", r.Name)
	fmt.Printf("Coordinates: (%d, %d)\n", r.X, r.Y)

	if r.IsStart {
		fmt.Println("Type: START")
	}

	if r.IsEnd {
		fmt.Println("Type: END")
	}
}
