package functions

import "strconv"

func ActiveBits(n int) int {
	bt := strconv.FormatInt(int64(n), 2)
	count := 0
	for _, r := range bt {
		if r == '1' {
			count++
		}
	}
	return count
}
