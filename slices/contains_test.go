package slices_test

import (
	"fmt"
	"testing"

	"github.com/jvcoutinho/lambda/slices"
	"github.com/stretchr/testify/require"
)

func ExampleContains() {
	students := []string{"Bob", "Alice"}

	fmt.Println(slices.Contains(students, "Bob"))
	fmt.Println(slices.Contains(students, "Camille"))
	// Output:
	// true
	// false
}

func TestContains(t *testing.T) {
	t.Parallel()

	type TestCase[T comparable] struct {
		description    string
		currentSlice   []T
		element        T
		expectedResult bool
	}

	tests := []TestCase[int]{
		{
			description:    "GivenSliceIsEmpty_ShouldReturnFalse",
			currentSlice:   []int{},
			element:        2,
			expectedResult: false,
		},
		{
			description:    "GivenSliceDoesNotContainElement_ShouldReturnFalse",
			currentSlice:   []int{2, 3},
			element:        4,
			expectedResult: false,
		},
		{
			description:    "GivenSliceContainsElement_ShouldReturnTrue",
			currentSlice:   []int{2, 3},
			element:        2,
			expectedResult: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()

			// Act
			actualResult := slices.Contains(test.currentSlice, test.element)

			// Assert
			require.Equal(t, test.expectedResult, actualResult)
		})
	}
}
