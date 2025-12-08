package piscine

func ShoppingListSort(slice []string) []string {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if len(slice[j]) < len(slice[i]) {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}
