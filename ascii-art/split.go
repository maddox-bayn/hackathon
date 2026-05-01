package main

import "strings"

func SplitChar(input string) []string {
	sliceInput := strings.Split(input, "\\n")
	return sliceInput
}
