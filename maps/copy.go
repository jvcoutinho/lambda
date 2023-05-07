package maps

// Copy copies all key-value pairs from m into a new map.
//
// If m is nil, Copy returns nil.
func Copy[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}

	result := make(map[K]V, len(m))
	for key, value := range m {
		result[key] = value
	}

	return result
}
