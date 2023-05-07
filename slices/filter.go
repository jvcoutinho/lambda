package slices

// Filter returns all the elements e in slice where predicate(e) equals true.
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(slice))

	for _, e := range slice {
		if predicate(e) {
			result = append(result, e)
		}
	}

	return result
}
