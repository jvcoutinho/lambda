package slices

// Contains returns true if and only if slice contains elem by the default equality operator (==).
func Contains[T comparable](slice []T, elem T) bool {
	for _, e := range slice {
		if e == elem {
			return true
		}
	}

	return false
}
