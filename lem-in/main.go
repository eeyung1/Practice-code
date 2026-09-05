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

	startRoom := ""
	endRoom := ""

	for _, room := range parsed.Rooms {
		if room.IsStart {
			startRoom = room.Name
		}

		if room.IsEnd {
			endRoom = room.Name
		}
	}

	blocked := make(map[string]bool)

	path := BFS(graph, startRoom, endRoom, blocked)

	fmt.Println("Path:", path)

	fmt.Println(graph)
}