package main

import "fmt"

func main() {
	parsed, err := ReadFile("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Ants:", parsed.Ants)
	fmt.Println("Rooms loaded:", len(parsed.Rooms))

	for _, room := range parsed.Rooms {
		room.Display()
	}
}