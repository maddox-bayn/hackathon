package main

import (
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) string {

	var b strings.Builder

	if input == "" {
		return ""
	}

	input = strings.ReplaceAll(input, "\n", "\\n")

	siceINput := SplitInput(input)

	isOnlynewline := true

	for _, word := range siceINput {
		if word != "" {
			isOnlynewline = false
		}
	}

	if isOnlynewline {
		for i := 0; i < len(siceINput)-1; i++ {
			b.WriteString("\n")
		}
		return b.String()
	}
	for i, word := range siceINput {
		validate(word)

		if word == "" {
			if i < len(siceINput)-1 {
				b.WriteString("\n")
			}
			continue
		}

		asciiTable := RenderLine(word, banner)

		for i := 0; i < len(asciiTable); i++ {
			b.WriteString(asciiTable[i] + "\n")
		}
	}
	return b.String()
}
