package main

func BFS(graph *Graph, start string, end string) []string {
	roomQueue := []string{}
	visited := make(map[string]bool)
	parent := make(map[string]string)

	roomQueue = append(roomQueue, start)
	visited[start] = true

	_ = parent

	return []string{}
}