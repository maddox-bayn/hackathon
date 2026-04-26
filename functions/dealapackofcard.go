package functions

import "fmt"

func DealAPackOfCards(deck []int) {
	cardNumber := 0

	for i := 1; i <= 4; i++ {
		a := deck[cardNumber]
		b := deck[cardNumber+1]
		c := deck[cardNumber+2]

		fmt.Printf("Plater %d: %d, %d, %d\n", i, a, b, c)

		cardNumber += 3
	}
}
