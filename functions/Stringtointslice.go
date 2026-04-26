package functions

func StringToIntSlice(str string) []int {
	var result []int
	for _, cha := range str {
		result = append(result, int(cha))
	}
	return result
}
