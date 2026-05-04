package functions

import (
	"cmp"
	"slices"
)

func ShoppingListSort(slice []string) []string {

	for i := 0; i < len(slice); i++ {
		for j := 0; j < i; j++ {

			if len(slice[i]) < len(slice[j]) {

				slice[i], slice[j] = slice[j], slice[i]

			}
		}
	}
	return slice
}

func ShoppingListSort1(slice []string) []string {
	slices.SortFunc(slice, func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	})
	return slice

}
