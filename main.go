package main

import "fmt"

func Unmatch(a []int) int {
	match := make(map[int]int, 0)
	for _, i := range a {
		match[i]++
	}
	for k := range match {
		if match[k] == 1 {
			return k
		}
	}
	// for _, num := range a {
	// 	if match[num] % 2 != 0 {
	// 		return num
	// 	}
	// }
	return -1
}

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 1, 2, 3, 4}
	fmt.Println(Unmatch(a))
}
