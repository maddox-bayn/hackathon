package functions

import "strings"

func Join(strs []string, sep string) string {
	var b strings.Builder
	for i := 0; i < len(strs); i++ {
		b.WriteString(strs[i])
		if i != len(strs)-1 {
			b.WriteString(sep)
		}
	}
	return b.String()
	//return strings.Join(strs, sep)
}
