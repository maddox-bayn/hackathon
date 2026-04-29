package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]

	asciiTable, _ := LoadBanner("standard.txt")

	fmt.Print(GenerateArt(input, asciiTable))

}
