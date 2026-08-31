package main


type Node struct {
	Next *Node
	Value interface{}
}

type List struct {
	Head *Node
	Tail *Node
}

func ListPushBack(l *List, value interface{}) {
	node := &Node{Value: value}
	if l.Tail == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
}