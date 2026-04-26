package main

import "fmt"

func DealAPackOfCards(deck []int) {
	cardIndex := 0
	//a, b, c := 0, 0, 0
	for i := 1; i <= 4; i++ {

		a := deck[cardIndex]
		b := deck[cardIndex+1]
		c := deck[cardIndex+2]

		cardIndex += 3

		fmt.Printf("Player %d: %d, %d, %d\n", i, a, b, c)
	}

}

func main() {
	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	DealAPackOfCards(deck)
}
