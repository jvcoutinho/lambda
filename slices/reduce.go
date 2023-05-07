package slices

// Reduce aggregates the elements of slice to a single value using aggregate.
//
// If slice is empty, Reduce returns seed.
func Reduce[T, R any](slice []T, seed R, aggregate func(currentResult R, currentElement T) R) R {
	result := seed

	for _, e := range slice {
		result = aggregate(result, e)
	}

	return result
}
