package utils

func Contains(slice []int, element int) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}

	return false
}