package functions

func JumpOver(str string) string {
	// runes := []rune(str)
	// var result []rune

	// for i := 0; i < 3; i++ {
	// 	if len(str) > 0 && len(runes) >= 3 {
	// 		// if i > len(runes)-1 {
	// 		// 	break
	// 		// }
	// 		result = append(result, runes[i])
	// 	}
	// }
	// return string(result) + "\n"

	if len(str) < 3 {
		return "\n"
	}
	return str[:3] + "\n"
}
