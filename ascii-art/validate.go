package main

import "fmt"

func validate(str string) (rune, error) {
	for _, c := range str {
		if c < 32 || c > 126 {
			return c, fmt.Errorf("Charater not in range %c", c)
		}
	}
	return 0, nil
}
