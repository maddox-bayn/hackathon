package main

func ListReverse(l *List) {

	if l.Head == nil || l.Tail == nil {
		return
	}

	oldHead := l.Head

	var prev *Node = nil
	current := l.Head
	// spotter to hold onto the next node
	var next *Node = nil

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
	l.Tail = oldHead
} 