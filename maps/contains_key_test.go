package maps_test

import (
	"fmt"
	"testing"

	"github.com/jvcoutinho/lambda/maps"
	"github.com/stretchr/testify/require"
)

func ExampleContainsKey() {
	age := map[string]int{
		"Bob":   25,
		"Alice": 19,
	}

	fmt.Println(maps.ContainsKey(age, "Bob"))
	fmt.Println(maps.ContainsKey(age, "Camille"))

	// Output:
	// true
	// false
}

func TestContainsKey(t *testing.T) {
	t.Parallel()

	type TestCase[K comparable, V any] struct {
		description    string
		currentMap     map[K]V
		key            K
		expectedResult bool
	}

	tests := []TestCase[string, int]{
		{
			description:    "GivenMapContainsKey_ShouldReturnTrue",
			currentMap:     map[string]int{"2": 2, "3": 3},
			key:            "2",
			expectedResult: true,
		},
		{
			description:    "GivenMapDoesNotContainKey_ShouldReturnFalse",
			currentMap:     map[string]int{"2": 2, "3": 3},
			key:            "5",
			expectedResult: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()

			// Act
			actualResult := maps.ContainsKey(test.currentMap, test.key)

			// Assert
			require.Equal(t, test.expectedResult, actualResult)
		})
	}
}
