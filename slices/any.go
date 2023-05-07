package slices

// Any returns true if and only if there's at least one element e in slice where predicate(e) equals true.
func Any[T any](slice []T, predicate func(T) bool) bool {
	for _, e := range slice {
		if predicate(e) {
			return true
		}
	}

	return false
}
