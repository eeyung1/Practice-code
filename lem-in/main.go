package main

import "fmt"

func main() {
	rooms, err := ReadFile("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Rooms loaded:", len(rooms))

	for _, room := range rooms {
		room.Display()
	}
}
