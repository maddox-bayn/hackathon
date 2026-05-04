package functions

import "fmt"

func CompareSlice(a []string) {
	for _, word := range a {
		if word == "01" || word == "galaxy" || word == "galaxy 01" {
			fmt.Println("Alert!!!")
		}
	}
}
