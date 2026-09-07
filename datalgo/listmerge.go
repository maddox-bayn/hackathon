package main

func ListMerge(l1 *List, l2 *List) {
	current := l2.Head
	for current != nil {
		l1.Tail.Next = current
		l1.Tail = current
		current = current.Next
	}

}