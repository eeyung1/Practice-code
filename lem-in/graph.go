package main

import "sort"

type Graph struct { Neighbors map[string][]string }

func BuildGraph(data *ParsedData) *Graph {
 graph:=&Graph{Neighbors:make(map[string][]string)}
 for _,room:=range data.Rooms {graph.Neighbors[room.Name]=nil}
 for _,tunnel:=range data.Tunnels {
  graph.Neighbors[tunnel.From]=append(graph.Neighbors[tunnel.From],tunnel.To)
  graph.Neighbors[tunnel.To]=append(graph.Neighbors[tunnel.To],tunnel.From)
 }
 for room:=range graph.Neighbors {sort.Strings(graph.Neighbors[room])}
 return graph
}
