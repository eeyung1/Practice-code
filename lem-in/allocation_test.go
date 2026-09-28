package main

import "testing"

func TestAllocateAntsBalancesThreePaths(t *testing.T) {
	paths := [][]string{
		{"S", "A", "E"},           // 2 tunnels
		{"S", "B", "C", "E"},      // 3 tunnels
		{"S", "D", "F", "G", "E"}, // 4 tunnels
	}

	got := AllocateAnts(paths, 6)
	want := []int{3, 2, 1}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("allocation = %v; want %v", got, want)
		}
	}
}

func TestAssignAntsUsesAllocation(t *testing.T) {
	paths := [][]string{
		{"S", "A", "E"},
		{"S", "B", "C", "E"},
		{"S", "D", "F", "G", "E"},
	}

	ants := AssignAnts(paths, 6)
	counts := make([]int, len(paths))

	for _, ant := range ants {
		for i, path := range paths {
			if len(ant.Path) == len(path) {
				counts[i]++
				break
			}
		}
	}

	want := []int{3, 2, 1}
	for i := range want {
		if counts[i] != want[i] {
			t.Fatalf("assigned ants = %v; want %v", counts, want)
		}
	}
}