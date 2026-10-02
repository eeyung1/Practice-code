package main

import (
	"math/rand"
	"testing"
)

func TestPathSelectionAgainstExhaustiveSearch(t *testing.T) {
	random := rand.New(rand.NewSource(17))
	for sample := 0; sample < 150; sample++ {
		names := []string{"S", "A", "B", "C", "D", "E"}
		data := &ParsedData{Ants: 1}
		for i, name := range names {
			data.Rooms = append(data.Rooms, Room{Name: name, X: i, IsStart: i == 0, IsEnd: i == len(names)-1})
		}
		for i := range names {
			for j := i + 1; j < len(names); j++ {
				if random.Intn(3) == 0 {
					data.Tunnels = append(data.Tunnels, Tunnel{From: names[i], To: names[j]})
				}
			}
		}
		graph := BuildGraph(data)
		all := [][]string{}
		var walk func(string, []string, map[string]bool)
		walk = func(room string, path []string, seen map[string]bool) {
			if room == "E" {
				all = append(all, append([]string(nil), path...))
				return
			}
			for _, next := range graph.Neighbors[room] {
				if !seen[next] {
					seen[next] = true
					walk(next, append(path, next), seen)
					delete(seen, next)
				}
			}
		}
		walk("S", []string{"S"}, map[string]bool{"S": true})
		for ants := 1; ants <= 6; ants++ {
			data.Ants = ants
			want := 0
			var choose func(int, [][]string, map[string]bool)
			choose = func(index int, set [][]string, used map[string]bool) {
				if len(set) > 0 {
					for turns := 1; turns <= ants+len(names); turns++ {
						capacity := 0
						for _, path := range set {
							slots := turns - (len(path) - 1) + 1
							if slots > 0 {
								capacity += slots
							}
						}
						if capacity >= ants {
							if want == 0 || turns < want {
								want = turns
							}
							break
						}
					}
				}
				for i := index; i < len(all); i++ {
					ok := true
					for _, room := range all[i][1 : len(all[i])-1] {
						if used[room] {
							ok = false
						}
					}
					if !ok {
						continue
					}
					for _, room := range all[i][1 : len(all[i])-1] {
						used[room] = true
					}
					choose(i+1, append(set, all[i]), used)
					for _, room := range all[i][1 : len(all[i])-1] {
						delete(used, room)
					}
				}
			}
			choose(0, nil, map[string]bool{})
			turns, err := Solve(data)
			if want == 0 {
				if err == nil {
					t.Fatalf("sample %d: accepted disconnected map", sample)
				}
				continue
			}
			if err != nil {
				t.Fatalf("sample %d: %v", sample, err)
			}
			if len(turns) != want {
				t.Fatalf("sample %d ants %d: got %d turns; exhaustive optimum %d; edges %+v", sample, ants, len(turns), want, data.Tunnels)
			}
			checkMoves(t, data, turns)
		}
	}
}

func TestFindPathsTerminatesOnDirectConnection(t *testing.T) {
	graph := &Graph{Neighbors: map[string][]string{"S": {"E"}, "E": {"S"}}}
	paths := FindPaths(graph, "S", "E")
	if len(paths) != 1 || len(paths[0]) != 2 {
		t.Fatalf("paths: %v", paths)
	}
	if paths := FindPaths(graph, "S", "S"); len(paths) != 0 {
		t.Fatalf("same endpoints: %v", paths)
	}
}

func TestSimulationRejectsBadPaths(t *testing.T) {
	cases := [][][]string{nil, {{"S"}}, {{"S", "A", "E"}, {"S", "A", "E"}}, {{"S", "E"}, {"S", "E"}}, {{"S", "A", "A", "E"}}, {{"S", "E"}, {"X", "E"}}}
	for i, paths := range cases {
		if _, err := SimulateTurns(paths, 3); err == nil {
			t.Errorf("case %d accepted", i)
		}
	}
}

func TestResidualFlowReroutesGreedyChoice(t *testing.T) {
 data := &ParsedData{Ants: 6}
 for i,name := range []string{"S","A","B","C","D","E"} {
  data.Rooms = append(data.Rooms, Room{Name:name, X:i, IsStart:name=="S", IsEnd:name=="E"})
 }
 data.Tunnels = []Tunnel{{"S","A"},{"A","B"},{"B","E"},{"A","C"},{"C","E"},{"S","D"},{"D","B"}}
 turns,err := Solve(data)
 if err != nil { t.Fatal(err) }
 if len(turns) != 5 { t.Fatalf("got %d turns; want 5",len(turns)) }
 checkMoves(t,data,turns)
}
