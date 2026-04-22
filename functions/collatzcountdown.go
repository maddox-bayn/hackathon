package functions

func CollatzCountdown(start int) int {

	step := 0
	if start <= 0 {
		return -1
	}
	for start >= 1 && start != 1 {
		if start%2 == 0 {
			start = start / 2
			step++
		} else {
			start = (start * 3) + 1
			step++
		}
	}
	return step
}
