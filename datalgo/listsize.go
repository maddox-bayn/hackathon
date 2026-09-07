package main

import "fmt"

func Listsize(l *List) int {
	count := 0
	for current := l.Head; current != nil; current = current.Next {
		count++
		fmt.Println(current.Data)
	}
	return count
}