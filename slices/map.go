package slices

// Map returns a new slice containing all elements of slice mapped by transform.
func Map[T, V any](slice []T, transform func(T) V) []V {
	result := make([]V, len(slice))

	for i := 0; i < len(slice); i++ {
		result[i] = transform(slice[i])
	}

	return result
}
