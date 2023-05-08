package set_test

import (
	"testing"

	"github.com/jvcoutinho/lambda/set"
	"github.com/stretchr/testify/require"
)

func TestSet_Add(t *testing.T) {
	t.Parallel()

	type TestCase[T comparable] struct {
		description      string
		initialElements  []T
		element          T
		expectedElements []T
	}

	tests := []TestCase[int]{
		{
			description:      "GivenSetDoesNotContainElement_ShouldAdd",
			initialElements:  []int{3},
			element:          2,
			expectedElements: []int{2, 3},
		},
		{
			description:      "GivenSetContainsElement_ShouldDoNothing",
			initialElements:  []int{2, 3},
			element:          2,
			expectedElements: []int{2, 3},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()

			// Arrange
			set := set.New(test.initialElements...)

			// Act
			set.Add(test.element)

			// Assert
			require.ElementsMatch(t, test.expectedElements, set.Enumerate())
		})
	}
}

func TestSet_Remove(t *testing.T) {
	t.Parallel()

	type TestCase[T comparable] struct {
		description      string
		initialElements  []T
		element          T
		expectedElements []T
	}

	tests := []TestCase[int]{
		{
			description:      "GivenSetIsEmpty_ShouldDoNothing",
			initialElements:  []int{},
			element:          2,
			expectedElements: []int{},
		},
		{
			description:      "GivenSetDoesNotContainElement_ShouldDoNothing",
			initialElements:  []int{3},
			element:          2,
			expectedElements: []int{3},
		},
		{
			description:      "GivenSetContainsElement_ShouldRemove",
			initialElements:  []int{2, 3},
			element:          2,
			expectedElements: []int{3},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()

			// Arrange
			set := set.New(test.initialElements...)

			// Act
			set.Remove(test.element)

			// Assert
			require.ElementsMatch(t, test.expectedElements, set.Enumerate())
		})
	}
}

func TestSet_Contains(t *testing.T) {
	t.Parallel()

	type TestCase[T comparable] struct {
		description     string
		initialElements []T
		element         T
		expectedResult  bool
	}

	tests := []TestCase[int]{
		{
			description:     "GivenSetIsEmpty_ShouldReturnFalse",
			initialElements: []int{},
			element:         2,
			expectedResult:  false,
		},
		{
			description:     "GivenSetDoesNotContainElement_ShouldReturnFalse",
			initialElements: []int{3},
			element:         2,
			expectedResult:  false,
		},
		{
			description:     "GivenSetContainsElement_ShouldReturnTrue",
			initialElements: []int{2, 3},
			element:         2,
			expectedResult:  true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()

			// Arrange
			set := set.New(test.initialElements...)

			// Act
			actualResult := set.Contains(test.element)

			// Assert
			require.Equal(t, test.expectedResult, actualResult)
		})
	}
}

func TestSet_Len(t *testing.T) {
	t.Parallel()

	type TestCase[T comparable] struct {
		description     string
		initialElements []T
	}

	tests := []TestCase[int]{
		{
			description:     "GivenSetIsEmpty_ShouldReturnZero",
			initialElements: []int{},
		},
		{
			description:     "ShouldReturnLength",
			initialElements: []int{3, 4, 5, 6},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()

			// Arrange
			set := set.New(test.initialElements...)

			// Act
			actualResult := set.Len()

			// Assert
			require.Len(t, test.initialElements, actualResult)
		})
	}
}

func TestSet_Enumerate(t *testing.T) {
	t.Parallel()

	type TestCase[T comparable] struct {
		description     string
		initialElements []T
	}

	tests := []TestCase[int]{
		{
			description:     "GivenSetIsEmpty_ShouldReturnEmptySlice",
			initialElements: []int{},
		},
		{
			description:     "ShouldReturnElements",
			initialElements: []int{3, 4, 5, 6},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()

			// Arrange
			set := set.New(test.initialElements...)

			// Act
			actualResult := set.Enumerate()

			// Assert
			require.ElementsMatch(t, test.initialElements, actualResult)
		})
	}
}
