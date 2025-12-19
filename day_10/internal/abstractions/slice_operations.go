package abstractions

func Contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func IndexOf(slice []int, value int) int {
	for index, v := range slice {
		if v == value {
			return index
		}
	}
	return NotFound
}

func Clear(slice []int) {
	for index, _ := range slice {
		slice[index] = -1
	}
}
