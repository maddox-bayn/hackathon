package functions

func DescendAppendRange(max, min int) []int {

	var result []int

	if max <= min {
		return result
	}

	for i := max; i > min; i-- {
		result = append(result, i)
	}
	return result
}
