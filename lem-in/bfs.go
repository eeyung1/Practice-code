package main

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
