package main

import "strings"

func SplitInput(text string) []string {
	result := strings.Split(text, "\\n")
	return result
}
