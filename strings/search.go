package strings

// sliceContains will take the string and check that it exists
// within the slice given and return boolean
func sliceContains(key string, slice []string) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}

	return false
}
