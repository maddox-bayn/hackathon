package functions

func Rot14(s string) string {
	runes := []rune(s)

	for i, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			runes[i] = (runes[i] - 'a' + 14) % 26
			runes[i] = runes[i] + 'a'
		}

		if ch >= 'A' && ch <= 'Z' {
			runes[i] = (runes[i] - 'A' + 14) % 26
			runes[i] = runes[i] + 'A'
		}
	}
	return string(runes)
}
