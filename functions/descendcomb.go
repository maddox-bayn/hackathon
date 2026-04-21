package functions

import "fmt"

func DescendComb() {
	for i := 99; i >= 0; i-- {
		for j := 98; j >= 1; j-- {
			fmt.Printf("%02d %02d, ", i, j)
		}
	}
}
