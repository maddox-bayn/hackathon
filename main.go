package main

import (
	"fmt"
	"hackathon/functions"
)

func main() {
	fmt.Println(functions.FoodDeliveryTime("burger"))
	fmt.Println(functions.FoodDeliveryTime("chs"))
	fmt.Println(functions.FoodDeliveryTime("nuggets"))
	fmt.Println(functions.FoodDeliveryTime("burger") + functions.FoodDeliveryTime("nuggets") + functions.FoodDeliveryTime("chips"))
}
