package slices

// Equal returns true if and only if s1 and s2 have the same length
// and all elements from s1 and s2 (in increasing index order) are equal.
func Equal[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
