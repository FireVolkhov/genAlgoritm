package core

func SliceContainsInt(slice []int, item int) bool {
	for _, itemInSlice := range slice {
		if itemInSlice == item {
			return true
		}
	}
	return false
}
