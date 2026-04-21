package functions

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
	return -1
}
