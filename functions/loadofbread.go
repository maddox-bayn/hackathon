package functions

import (
	"strings"
)

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	var b strings.Builder
	count := 0
	for _, r := range str {
		if count == 5 {
			b.WriteRune(' ')
			count = 0
			continue
		}
		if r == ' ' {
			continue
		}
		b.WriteRune(r)
		count++
	}
	return b.String() + "\n"
}
