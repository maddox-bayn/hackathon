package main

func sortlist(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	var sorted *NodeI = nil

	current := l
	for current.Next != nil {
		sorted = sortInsert(sorted, current)
		current = current.Next
	}
	return sorted
}

func sortInsert(sortedHead *NodeI, newnode *NodeI) *NodeI {
	if sortedHead == nil || sortedHead.Data >= newnode.Data {
		newnode.Next = sortedHead
		return newnode
	}

	current := sortedHead
	for current.Next != nil && sortedHead.Data < newnode.Data {
		current = current.Next
	}
	newnode.Next = current.Next
	current.Next = newnode
	return sortedHead
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}
	if l == nil || l.Data >= data_ref {
		newNode.Next = l
		return newNode
	}

	current := l
	for current.Next != nil && current.Next.Data < newNode.Data {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
	return l
}
