package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	if l.Head == nil {
		// List is empty
		l.Head = newNode
		l.Tail = newNode
	} else {
		// List has at least one element
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
