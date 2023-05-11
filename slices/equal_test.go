package slices_test

import (
	"fmt"
	"testing"

	"github.com/jvcoutinho/lambda/slices"
	"github.com/stretchr/testify/require"
)

func ExampleEqual() {
	studentsInTennis := []string{"Bob", "Alice"}
	studentsInChess := []string{"Bob", "Alice"}
	studentsInVolley := []string{"Bob", "Camille"}

	fmt.Println(slices.Equal(studentsInTennis, studentsInChess))
	fmt.Println(slices.Equal(studentsInTennis, studentsInVolley))

	// Output:
	// true
	// false
}

func TestEqual(t *testing.T) {
	t.Parallel()

	type TestCase[T comparable] struct {
		description    string
		slice1         []T
		slice2         []T
		expectedResult bool
	}

	tests := []TestCase[int]{
		{
			description:    "GivenBothSlicesAreEmpty_ShouldReturnTrue",
			slice1:         []int{},
			slice2:         []int{},
			expectedResult: true,
		},
		{
			description:    "GivenSlice1IsLargerThanSlice2_ShouldReturnFalse",
			slice1:         []int{2, 3},
			slice2:         []int{2},
			expectedResult: false,
		},
		{
			description:    "GivenSlice2IsLargerThanSlice1_ShouldReturnFalse",
			slice1:         []int{2},
			slice2:         []int{2, 3},
			expectedResult: false,
		},
		{
			description:    "GivenSlicesDifferAtSomeIndex_ShouldReturnFalse",
			slice1:         []int{2, 5, 6},
			slice2:         []int{2, 3, 6},
			expectedResult: false,
		},
		{
			description:    "GivenSlicesDifferInOrder_ShouldReturnFalse",
			slice1:         []int{2, 6, 3},
			slice2:         []int{2, 3, 6},
			expectedResult: false,
		},
		{
			description:    "GivenSlicesDifferAtNoIndex_ShouldReturnTrue",
			slice1:         []int{2, 3, 6, 9},
			slice2:         []int{2, 3, 6, 9},
			expectedResult: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()

			// Act
			actualResult := slices.Equal(test.slice1, test.slice2)

			// Assert
			require.Equal(t, test.expectedResult, actualResult)
		})
	}
}
