// Package set contains an implementation for a set.
package set

// Set is a unordered sequence of unique elements.
type Set[T comparable] map[T]bool

// New creates a new Set.
func New[T comparable](initialElements ...T) Set[T] {
	set := Set[T](make(map[T]bool))

	for _, element := range initialElements {
		set.Add(element)
	}

	return set
}

// Add adds elem to this set.
func (s Set[T]) Add(elem T) {
	s[elem] = true
}

// Remove removes elem from this set.
//
// If this set does not contain elem, Remove does nothing.
func (s Set[T]) Remove(elem T) {
	delete(s, elem)
}

// Contains returns true if and only if this set contains elem.
func (s Set[T]) Contains(elem T) bool {
	_, ok := s[elem]

	return ok
}

// Len is the number of elements this set contains.
func (s Set[T]) Len() int {
	return len(s)
}

// Enumerate returns an ordered sequence of the elements of this set.
// The order of the elements is not deterministic.
func (s Set[T]) Enumerate() []T {
	result := make([]T, 0, s.Len())

	for e := range s {
		result = append(result, e)
	}

	return result
}
