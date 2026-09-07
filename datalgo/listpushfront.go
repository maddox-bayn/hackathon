package main

func ListPushFront(l *List, value interface{}) {
	newnode := &Node{Data: value}
	newnode.Next = l.Head
	l.Head = newnode
	if l.Tail == nil {
		l.Tail = newnode
	}
}