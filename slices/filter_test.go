package slices_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jvcoutinho/lambda/slices"
	"github.com/stretchr/testify/require"
)

func ExampleFilter() {
	students := []string{"Bob", "Alice", "Camille"}

	startsWithUpperB := func(s string) bool { return strings.HasPrefix(s, "B") }
	endsWithLowerA := func(s string) bool { return strings.HasSuffix(s, "a") }

	fmt.Println(slices.Filter(students, startsWithUpperB))
	fmt.Println(slices.Filter(students, endsWithLowerA))

	// Output:
	// [Bob]
	// []
}

func TestFilter(t *testing.T) {
	t.Parallel()

	type TestCase[T any] struct {
		description    string
		currentSlice   []T
		predicate      func(T) bool
		expectedResult []T
	}

	tests := []TestCase[int]{
		{
			description:    "GivenSliceIsEmpty_ShouldReturnEmptySlice",
			currentSlice:   []int{},
			predicate:      func(i int) bool { return i%2 == 0 },
			expectedResult: []int{},
		},
		{
			description:    "GivenPredicateReturnsTrueForNoElements_ShouldReturnEmptySlice",
			currentSlice:   []int{3, 5},
			predicate:      func(i int) bool { return i%2 == 0 },
			expectedResult: []int{},
		},
		{
			description:    "GivenPredicateReturnsTrueForSomeElements_ShouldReturnSliceWithThoseElements",
			currentSlice:   []int{2, 3, 4},
			predicate:      func(i int) bool { return i%2 == 0 },
			expectedResult: []int{2, 4},
		},
		{
			description:    "GivenPredicateReturnsTrueForAllElements_ShouldReturnSliceWithAllElements",
			currentSlice:   []int{2, 4, 6},
			predicate:      func(i int) bool { return i%2 == 0 },
			expectedResult: []int{2, 4, 6},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()

			// Act
			actualResult := slices.Filter(test.currentSlice, test.predicate)

			// Assert
			require.Equal(t, test.expectedResult, actualResult)
		})
	}
}
