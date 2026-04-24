package functions

import "fmt"

func Abort(a, b, c, d int) int {
	nums := []int{a, b, c, d}
	len := len(nums)
	idx := 0

	//	start := nums[0]
	for i, j := 0, 1; i < j && j < len; i, j = i+1, j+1 {
		if nums[i] > nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if len%2 != 0 {
		idx = (len + 1) / 2
		idx -= 1
	} else {
		idx = len / 2
	}
	fmt.Println(nums)
	return nums[idx]
}
