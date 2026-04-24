package functions


// worst approach
// func Compact(ptr *[]string) int {
// 	for i, v := range *ptr {
// 		if v == "" && i+1 <= len(*ptr) {
// 			*ptr = append((*ptr)[:i], (*ptr)[i+1:]...)
// 		}
// 	}
// 	return len(*ptr)
// }


// a better approach
func Compact(ptr *[]string) int {

	if ptr == nil {
		return 0
	}
	s := *ptr

	nonZeroIndex := 0

	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			s[nonZeroIndex] = s[i]
			nonZeroIndex++
		}
	}
	*ptr = s[:nonZeroIndex]

	return nonZeroIndex
}