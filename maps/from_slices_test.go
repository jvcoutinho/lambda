package maps_test

import (
	"fmt"
	"testing"

	"github.com/jvcoutinho/lambda/maps"
	"github.com/stretchr/testify/require"
)

func ExampleFromSlices() {
	students := []string{"Bob", "Alice"}
	ages := []int{25, 19}
	classes := []string{"A"}

	fmt.Println(maps.FromSlices(students, ages))
	fmt.Println(maps.FromSlices(students, classes))

	// Output:
	// map[Alice:19 Bob:25]
	// map[Bob:A]
}

func TestFromSlices(t *testing.T) {
	t.Parallel()

	type TestCase[K comparable, V any] struct {
		description    string
		keys           []K
		values         []V
		expectedResult map[K]V
	}

	tests := []TestCase[string, int]{
		{
			description:    "GivenBothSlicesAreEmpty_ShouldReturnEmptyMap",
			keys:           []string{},
			values:         []int{},
			expectedResult: map[string]int{},
		},
		{
			description:    "GivenKeysSliceIsEmpty_ShouldReturnEmptyMap",
			keys:           []string{},
			values:         []int{2, 3},
			expectedResult: map[string]int{},
		},
		{
			description:    "GivenValuesSliceIsEmpty_ShouldReturnEmptyMap",
			keys:           []string{"2", "3"},
			values:         []int{},
			expectedResult: map[string]int{},
		},
		{
			description:    "GivenKeysSliceIsLargerThanValuesSlice_ShouldReturnMapWithLengthOfValuesSlice",
			keys:           []string{"2", "3", "4"},
			values:         []int{2, 3},
			expectedResult: map[string]int{"2": 2, "3": 3},
		},
		{
			description:    "GivenValuesSliceIsLargerThanKeysSlice_ShouldReturnMapWithLengthOfKeysSlice",
			keys:           []string{"2", "3"},
			values:         []int{2, 3, 4},
			expectedResult: map[string]int{"2": 2, "3": 3},
		},
		{
			description:    "GivenSlicesHaveTheSameLength_ShouldReturnMapWithAllElementsFromBothSlices",
			keys:           []string{"2", "3", "4"},
			values:         []int{2, 3, 4},
			expectedResult: map[string]int{"2": 2, "3": 3, "4": 4},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()

			// Act
			actualResult := maps.FromSlices(test.keys, test.values)

			// Assert
			require.Equal(t, test.expectedResult, actualResult)
		})
	}
}
