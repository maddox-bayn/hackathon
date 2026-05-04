package functions

func ReverseMenuIndex(menu []string) []string {
	for i, j := 0, len(menu)-1; i < j; i, j = i+1, j-1 {
		menu[i], menu[j] = menu[j], menu[i]
	}
	return menu
}