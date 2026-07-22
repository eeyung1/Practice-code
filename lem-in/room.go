package main

import "fmt"

type Room struct {
	Name string
	X    int
	Y    int
}

func (r Room) Display() {
	fmt.Printf("Room: %s\n", r.Name)
	fmt.Printf("Coordinates: (%d, %d)\n", r.X, r.Y)
}
