package slices_test

import (
	"fmt"
	"testing"

	"github.com/jvcoutinho/lambda/slices"
	"github.com/stretchr/testify/require"
)

func ExampleElementAtOrDefault() {
	students := []string{"Bob", "Alice"}

	fmt.Println(slices.ElementAtOrDefault(students, 0))
	fmt.Println(slices.ElementAtOrDefault(students, 5))
	// Output:
	// Bob
	//
}

func TestElementAtOrDefault(t *testing.T) {
	t.Parallel()

	type TestCase[T any] struct {
		description    string
		slice          []T
		index          int
		expectedResult T
	}

	tests := []TestCase[int]{
		{
			description:    "GivenIndexIsLessThanZero_ShouldReturnDefaultValue",
			slice:          []int{2, 3},
			index:          -2,
			expectedResult: 0,
		},
		{
			description:    "GivenIndexIsGreaterThanSliceLength_ShouldReturnDefaultValue",
			slice:          []int{2, 3},
			index:          4,
			expectedResult: 0,
		},
		{
			description:    "GivenIndexIsInSliceBounds_ShouldReturnElement",
			slice:          []int{2, 3},
			index:          1,
			expectedResult: 3,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()

			// Act
			actualResult := slices.ElementAtOrDefault(test.slice, test.index)

			// Assert
			require.Equal(t, test.expectedResult, actualResult)
		})
	}
}
