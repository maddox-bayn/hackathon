package main

import "strings"

func RenderLine(text string, banner map[rune][]string) []string {
	var result []string
	var b strings.Builder

	for i := 0; i < 8; i++ {
		for _, c := range text {
			if line, ok := banner[c]; ok {
				b.WriteString(line[i])
			}
		}
		result = append(result, b.String())
		b.Reset()
	}
	return result
}
