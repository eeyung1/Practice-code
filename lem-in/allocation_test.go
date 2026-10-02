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

	if len(got) != len(want) { t.Fatalf("allocation size = %d; want %d", len(got), len(want)) }
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
	if len(ants) != 6 { t.Fatalf("created %d ants; want 6", len(ants)) }
	for i,ant := range ants { if ant.ID != i+1 || ant.Position != 0 { t.Fatalf("invalid initial ant: %+v", ant) } }
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
func TestAllocationEdgeCases(t *testing.T) {
 cases := []struct { paths [][]string; ants int; want []int }{
  {nil, 3, []int{}},
  {[][]string{{"S", "A", "E"}}, 5, []int{5}},
  {[][]string{{"S", "A", "E"}, {"S", "B", "C", "D", "E"}}, 1, []int{1,0}},
  {[][]string{nil, {"S"}, {"S", "E"}}, 2, []int{0,0,2}},
  {[][]string{{"S", "E"}}, 0, []int{0}},
 }
 for _,tc := range cases {
  got:=AllocateAnts(tc.paths,tc.ants)
  if len(got)!=len(tc.want) {t.Fatalf("allocation length %d; want %d",len(got),len(tc.want))}
  for i:=range got {if got[i]!=tc.want[i] {t.Fatalf("allocation %v; want %v",got,tc.want)}}
 }
}
