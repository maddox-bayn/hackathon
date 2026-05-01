package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(fileName string) (map[rune][]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Error.. with ReadFile() %e", err)
	}

	fileLines := strings.Split(string(data), "\n")

	if len(fileLines) < 855 {
		return nil, fmt.Errorf("Erroor... with banner file len")
	}
	asciiTable := make(map[rune][]string)

	for i := 32; i < 127; i++ {
		startidx := (i-32)*9 + 1
		endidx := startidx + 8

		asciiTable[rune(i)] = fileLines[startidx:endidx]
	}
	return asciiTable, nil
}
