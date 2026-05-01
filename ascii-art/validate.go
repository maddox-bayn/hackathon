package main

import "fmt"

func ValidateInput(str string) (rune, error) {
	for _, char := range str {
		if char < 32 || char > 126 {
			return char, fmt.Errorf("charater not in range %c", char)
		}
	}
	return 0, nil
}
