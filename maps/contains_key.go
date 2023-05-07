package maps

// ContainsKey returns true if and only if elem is a key in m.
func ContainsKey[K comparable, V any](m map[K]V, elem K) bool {
	_, ok := m[elem]

	return ok
}
