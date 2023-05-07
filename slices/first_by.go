package slices

// FirstBy returns the first element e in slice where predicate(e) equals true.
//
// If predicate returns false for all elements in slice or if slice is empty, FirstBy returns ok = false.
func FirstBy[T any](slice []T, predicate func(T) bool) (result T, ok bool) {
	for _, e := range slice {
		if predicate(e) {
			result, ok = e, true

			break
		}
	}

	return
}
