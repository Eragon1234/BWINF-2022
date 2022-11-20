package utils

// IndexOf searches the slice s for element e and returns its index.
// returns -1 if element does not exist
func IndexOf[T comparable](s []T, e T) int {
	for i, t := range s {
		if t == e {
			return i
		}
	}
	return -1
}
