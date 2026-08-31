package main

import "fmt"

func main() {
	parsed, err := ReadFile("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	graph := BuildGraph(parsed)

	fmt.Println("Ants:", parsed.Ants)
	fmt.Println("Rooms loaded:", len(parsed.Rooms))

	for _, room := range parsed.Rooms {
		room.Display()
	}

	fmt.Println("Tunnels loaded:", len(parsed.Tunnels))

	for _, tunnel := range parsed.Tunnels {
		fmt.Printf("Tunnel: %s - %s\n", tunnel.From, tunnel.To)
	}

	fmt.Println(graph)
}
