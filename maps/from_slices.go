package maps

// FromSlices builds a map from two slices (keys and values) in increasing index order.
//
// If a slice has more elements than the other, FromSlices will compute only at the minimum length between the two.
func FromSlices[K comparable, V any](keys []K, values []V) map[K]V {
	result := make(map[K]V)

	minLength := len(keys)
	if minLength > len(values) {
		minLength = len(values)
	}

	for i := 0; i < minLength; i++ {
		result[keys[i]] = values[i]
	}

	return result
}
