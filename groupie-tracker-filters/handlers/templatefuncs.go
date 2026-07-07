package handlers

func ContainsMember(
	selected []int,
	member int,
) bool {
	for _, value := range selected {
		if value == member {
			return true
		}
	}

	return false
}