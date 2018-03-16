package twitter

// isStringInSlice returns true if the string s is in slice of strings
// slice, otherwise returns false.
func isStringInSlice(s string, slice []string) bool {
	for _, i := range slice {
		if s == i {
			return true
		}
	}
	return false
}
