package piscine

func ListSize(l *List) int {
	count := 0
	for current := l.Head; current != nil; current = current.Next {
		count++
	}
	return count
	
}