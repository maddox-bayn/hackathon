package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]
	//fmt.Println(input)
	banner, _ := LoadBanner("standard.txt")

	fmt.Print(GenerateArt(input, banner))
}
