package main

import "fmt"

type Graph struct {
	Neighbors map[string][]string
}

func BuildGraph(data *ParsedData) *Graph {
	graph := &Graph{
		Neighbors: make(map[string][]string),
	}

	for _, tunnel := range data.Tunnels {
		graph.Neighbors[tunnel.From] =
			append(graph.Neighbors[tunnel.From], tunnel.To)

		graph.Neighbors[tunnel.To] =
			append(graph.Neighbors[tunnel.To], tunnel.From)
	}

	fmt.Println("Graph:")

	for room, neighbors := range graph.Neighbors {
		fmt.Println(room, "->", neighbors)
	}

	return graph
}
