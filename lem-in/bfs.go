package main

import "fmt"

type Ant struct {
	ID       int
	Path     []string
	Position int
}

func BFS(graph *Graph, start string, end string, blocked map[string]bool) []string {
	if start == end {
		return []string{start}
	}

	roomQueue := []string{}
	visited := make(map[string]bool)
	parent := make(map[string]string)

	roomQueue = append(roomQueue, start)
	visited[start] = true

	for len(roomQueue) > 0 {
		current := roomQueue[0]
		roomQueue = roomQueue[1:]

		for _, neighbor := range graph.Neighbors[current] {
			if blocked[neighbor] {
				continue
			}

			if visited[neighbor] {
				continue
			}

			visited[neighbor] = true
			parent[neighbor] = current
			roomQueue = append(roomQueue, neighbor)

			if neighbor == end {
				path := []string{}
				currentRoom := end

				for currentRoom != start {
					path = append(path, currentRoom)
					currentRoom = parent[currentRoom]
				}

				path = append(path, start)

				for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
					path[i], path[j] = path[j], path[i]
				}

				return path
			}
		}
	}

	return []string{}
}

func FindPaths(graph *Graph, start string, end string) [][]string {
	paths := [][]string{}

	blocked := make(map[string]bool)

	for {
		path := BFS(graph, start, end, blocked)

		if len(path) == 0 {
			break
		}

		paths = append(paths, path)

		for i := 1; i < len(path)-1; i++ {
			blocked[path[i]] = true
		}
	}

	return paths
}

func Simulate(paths [][]string, antCount int) {
	ants := []Ant{}

	for i := 0; i < antCount; i++ {
		path := paths[i%len(paths)]

		ants = append(ants, Ant{
			ID:       i + 1,
			Path:     path,
			Position: 0,
		})
	}

	finished := 0
	turn := 0

	for finished < antCount {
		turn++

		occupied := make(map[string]bool)
		moves := []string{}

		for i := range ants {
			ant := &ants[i]

			if ant.Position == len(ant.Path)-1 {
				continue
			}

			nextRoom := ant.Path[ant.Position+1]

			if nextRoom != ant.Path[len(ant.Path)-1] && occupied[nextRoom] {
				continue
			}

			ant.Position++

			if nextRoom != ant.Path[len(ant.Path)-1] {
				occupied[nextRoom] = true
			}

			if ant.Position == len(ant.Path)-1 {
				finished++
			}

			moves = append(moves,
				fmt.Sprintf("L%d-%s", ant.ID, nextRoom),
			)
		}

		fmt.Println("Turn", turn, ":", moves)

	}
}
