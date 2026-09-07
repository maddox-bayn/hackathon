package main

func ListAt(l *Node, pos int) *Node{
	current := l
	for i := 0; i < pos; i++ {
		if current == nil {
			return nil
		}
		current = current.Next
	}
	return current
}

