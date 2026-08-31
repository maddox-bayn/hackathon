package main

func Listlast(l *List) interface{} {
	if l.Tail == nil {
		return nil
	}
	return l.Tail.Value
}
