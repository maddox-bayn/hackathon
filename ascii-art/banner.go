package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error with Readfile() = %e", err)
	}

	fileLines := strings.Split(string(data), "\n")

	if len(fileLines) < 855 {
		return  nil, fmt.Errorf("Error with banner file")
	}

	asciiTable := make(map[rune][]string)

	for i := 32; i < 127; i++ {
		startidx := (i-32) *9+1
		endIdx := startidx + 8

		asciiTable[rune(i)] = fileLines[startidx:endIdx]
	}

	return asciiTable, nil
}
