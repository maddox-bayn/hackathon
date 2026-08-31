package main

func ListPushFront(l *List, value interface{}) {
	newnode := &Node{Value: value}
	newnode.Next = l.Head
	l.Head = newnode
	if l.Tail == nil {
		l.Tail = newnode
	}
}