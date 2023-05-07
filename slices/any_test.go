package slices_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jvcoutinho/lambda/slices"
	"github.com/stretchr/testify/require"
)

func ExampleAny() {
	students := []string{"Bob", "Alice", "Camille"}

	startsWithUpperB := func(s string) bool { return strings.HasPrefix(s, "B") }
	endsWithLowerA := func(s string) bool { return strings.HasSuffix(s, "a") }

	fmt.Println(slices.Any(students, startsWithUpperB))
	fmt.Println(slices.Any(students, endsWithLowerA))

	// Output:
	// true
	// false
}

func TestAny(t *testing.T) {
	t.Parallel()

	type TestCase[T any] struct {
		description    string
		currentSlice   []T
		predicate      func(T) bool
		expectedResult bool
	}

	tests := []TestCase[int]{
		{
			description:    "GivenSliceIsEmpty_ShouldReturnFalse",
			currentSlice:   []int{},
			predicate:      func(i int) bool { return i%2 == 0 },
			expectedResult: false,
		},
		{
			description:    "GivenPredicateReturnsTrueForNoElement_ShouldReturnFalse",
			currentSlice:   []int{3, 5, 7},
			predicate:      func(i int) bool { return i%2 == 0 },
			expectedResult: false,
		},
		{
			description:    "GivenPredicateReturnsTrueForOneElement_ShouldReturnTrue",
			currentSlice:   []int{2, 3, 5, 7},
			predicate:      func(i int) bool { return i%2 == 0 },
			expectedResult: true,
		},
		{
			description:    "GivenPredicateReturnsTrueForAllElements_ShouldReturnTrue",
			currentSlice:   []int{2, 4, 6, 8},
			predicate:      func(i int) bool { return i%2 == 0 },
			expectedResult: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()

			// Act
			actualResult := slices.Any(test.currentSlice, test.predicate)

			// Assert
			require.Equal(t, test.expectedResult, actualResult)
		})
	}
}
