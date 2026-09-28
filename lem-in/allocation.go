package main

// AllocateAnts returns the number of ants assigned to each path.
func AllocateAnts(paths [][]string, antCount int) []int {
	counts := make([]int, len(paths))

	if len(paths) == 0 {
		return counts
	}

	for ant := 0; ant < antCount; ant++ {
		best := 0
		bestFinish := len(paths[0]) - 1 + counts[0]

		for i := 1; i < len(paths); i++ {
			projectedFinish := len(paths[i]) - 1 + counts[i]

			if projectedFinish < bestFinish {
				best = i
				bestFinish = projectedFinish
			}
		}

		counts[best]++
	}

	return counts
}

func AssignAnts(paths [][]string, antCount int) []Ant {
	counts := AllocateAnts(paths, antCount)
	ants := []Ant{}
	nextID := 1

	for pathIndex, count := range counts {
		for assigned := 0; assigned < count; assigned++ {
			ants = append(ants, Ant{
				ID:       nextID,
				Path:     paths[pathIndex],
				Position: 0,
			})
			nextID++
		}
	}

	return ants
}