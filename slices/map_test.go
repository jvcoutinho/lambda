package slices_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/jvcoutinho/lambda/slices"
	"github.com/stretchr/testify/require"
)

func ExampleMap() {
	students := []string{"Bob", "Alice", "Camille"}

	length := func(str string) int { return len(str) }

	fmt.Println(slices.Map(students, strings.ToLower))
	fmt.Println(slices.Map(students, length))

	// Output:
	// [bob alice camille]
	// [3 5 7]
}

func TestMap(t *testing.T) {
	t.Parallel()

	type TestCase[T, V any] struct {
		description    string
		currentSlice   []T
		transform      func(T) V
		expectedResult []V
	}

	tests := []TestCase[int, string]{
		{
			description:    "GivenSliceIsEmpty_ShouldReturnEmptySlice",
			currentSlice:   []int{},
			transform:      strconv.Itoa,
			expectedResult: []string{},
		},
		{
			description:    "ShouldReturnSliceWithTransformedElements",
			currentSlice:   []int{3, 5},
			transform:      strconv.Itoa,
			expectedResult: []string{"3", "5"},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()

			// Act
			actualResult := slices.Map(test.currentSlice, test.transform)

			// Assert
			require.Equal(t, test.expectedResult, actualResult)
		})
	}
}
