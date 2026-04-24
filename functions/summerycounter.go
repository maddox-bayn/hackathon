package functions

import "strings"

func ShoppingSummaryCounter(str string) map[string]int {
	list := make(map[string]int)
	sl := strings.Fields(str)
	for _, grocery := range sl {
		list[grocery]++
	}
	return list
}
