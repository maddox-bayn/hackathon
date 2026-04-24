package main

import (
	"fmt"
	"hackathon/functions"
)

const N = 6

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"
	fmt.Printf("%#v\n", a)
	for _, v := range a {
		fmt.Printf("%#v\n", v)
	}

	fmt.Println("Size after compacting:", functions.Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}
