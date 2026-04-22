package functions

type food struct {
	//order string
	preptime int
}

func FoodDeliveryTime(order string) int {
	nuggets := food{preptime: 12}
	chips := food{preptime: 10}
	burger := food{preptime: 15}
	menu := map[string]food{"nuggets": nuggets, "burger": burger, "chips": chips}

	result := menu[order].preptime

	if result == 0 {
		return 404
	}
	return result

}
