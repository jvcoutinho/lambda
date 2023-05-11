package slices_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/jvcoutinho/lambda/slices"
	"github.com/stretchr/testify/require"
)

func ExampleReduce() {
	ages := []int{25, 19, 31}

	max := func(max int, n int) int {
		if n > max {
			return n
		}

		return max
	}

	fmt.Println(slices.Reduce(ages, math.MinInt, max))
	// Output:
	// 31
}

func TestReduce(t *testing.T) {
	t.Parallel()

	type TestCase[T, R any] struct {
		description    string
		currentSlice   []T
		seed           R
		aggregate      func(R, T) R
		expectedResult R
	}

	tests := []TestCase[int, string]{
		{
			description:    "GivenSliceIsEmpty_ShouldReturnSeed",
			currentSlice:   []int{},
			seed:           "",
			aggregate:      func(concat string, n int) string { return concat + strconv.Itoa(n) },
			expectedResult: "",
		},
		{
			description:    "ShouldReturnAggregation",
			currentSlice:   []int{3, 5, 9},
			aggregate:      func(concat string, n int) string { return concat + strconv.Itoa(n) },
			expectedResult: "359",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()

			// Act
			actualResult := slices.Reduce(test.currentSlice, test.seed, test.aggregate)

			// Assert
			require.Equal(t, test.expectedResult, actualResult)
		})
	}
}
