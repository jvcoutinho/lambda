package slices_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jvcoutinho/lambda/slices"
	"github.com/stretchr/testify/require"
)

func ExampleFirstBy() {
	students := []string{"Bob", "Alice", "Camille"}

	startsWithUpperB := func(s string) bool { return strings.HasPrefix(s, "B") }
	endsWithLowerA := func(s string) bool { return strings.HasSuffix(s, "a") }
	endsWithLowerE := func(s string) bool { return strings.HasSuffix(s, "e") }

	fmt.Println(slices.FirstBy(students, startsWithUpperB))
	fmt.Println(slices.FirstBy(students, endsWithLowerA))
	fmt.Println(slices.FirstBy(students, endsWithLowerE))

	// Output:
	// Bob true
	//  false
	// Alice true
}

func TestFirstBy(t *testing.T) {
	t.Parallel()

	type TestCase[T comparable] struct {
		description    string
		currentSlice   []T
		predicate      func(T) bool
		expectedResult T
		expectedOk     bool
	}

	tests := []TestCase[int]{
		{
			description:    "GivenSliceIsEmpty_ShouldReturnFalse",
			currentSlice:   []int{},
			predicate:      func(i int) bool { return i%2 == 0 },
			expectedResult: 0,
			expectedOk:     false,
		},
		{
			description:    "GivenPredicateReturnsTrueForNoElement_ShouldReturnFalse",
			currentSlice:   []int{3, 5},
			predicate:      func(i int) bool { return i%2 == 0 },
			expectedResult: 0,
			expectedOk:     false,
		},
		{
			description:    "GivenPredicateReturnsTrueForOneElement_ShouldReturnIt",
			currentSlice:   []int{3, 4},
			predicate:      func(i int) bool { return i%2 == 0 },
			expectedResult: 4,
			expectedOk:     true,
		},
		{
			description:    "GivenPredicateReturnsTrueForSomeElements_ShouldReturnTheFirst",
			currentSlice:   []int{3, 4, 5, 6},
			predicate:      func(i int) bool { return i%2 == 0 },
			expectedResult: 4,
			expectedOk:     true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()

			// Act
			actualResult, actualOk := slices.FirstBy(test.currentSlice, test.predicate)

			// Assert
			require.Equal(t, test.expectedResult, actualResult)
			require.Equal(t, test.expectedOk, actualOk)
		})
	}
}
