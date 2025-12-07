package piscine

func Compact(ptr *[]string) int {
	list := *ptr
	count := 0

	for _, r := range list { // FOR each item in the list:
		if r != "" { // IF item is NOT empty:
			list[count] = r // put item at position count
			count = count + 1
		} // END IF
	} // END FOR

	*ptr = list[:count] // cut the list so it keeps only the first 'count' items

	return count // RETURN count
}
