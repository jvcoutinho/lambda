package slices

// ElementAtOrDefault returns the element of slice pointed by index.
//
// If index is less than zero or greater or equal than the length of the slice, ElementAtOrDefault
// returns the default value of type T.
func ElementAtOrDefault[T any](slice []T, index int) (result T) {
	if index >= 0 && index < len(slice) {
		result = slice[index]
	}

	return result
}
