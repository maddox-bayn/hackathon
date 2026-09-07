package main

type NodeI struct {
	Data int
	Next *NodeI
}


func ListSort(l *NodeI) *NodeI {
	// if node is empty 
	if l == nil || l.Next == nil {
		return l
	}
	
	// create sorted singly linked list
	var sorted *NodeI = nil
	
	current := l

	for current != nil {
		next := current.Next

		sorted = sortedInsert(sorted, current)

		current = next
	}
	return sorted
}

// helper function to insert a node into its correct sorted position
func sortedInsert(sortedHead *NodeI, newNode *NodeI) *NodeI {
	// to position the smallest number first or if it empty
	if sortedHead ==  nil || sortedHead.Data > newNode.Data {
		newNode.Next = sortedHead
		return newNode
	}

	// get the position to insert the right node
	current := sortedHead
	for current != nil && current.Next.Data < newNode.Data {
		current = current.Next
	}
	
	newNode.Next = current.Next
	current.Next = newNode

	return sortedHead
}