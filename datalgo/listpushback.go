package main


func ListPushBack(l *List, value interface{}) {
	node := &Node{Data: value}
	if l.Tail == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
}