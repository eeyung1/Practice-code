package piscine

func ListReverse(l *List) {
	if l.Head == nil || l.Head.Next == nil {
		return
	}

	var prev *NodeL
	current := l.Head
	next := current.Next

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Tail = l.Head
	l.Head = prev
}
