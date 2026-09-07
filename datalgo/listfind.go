package main

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	current := l.Head

	for current != nil {
		// Use the provided comparison function
		if comp(current.Data, ref) {
			// Return the address (&) of the Data field
			return &current.Data
		}
		current = current.Next
	}

	// Return nil if no matching element is found
	return nil
}
