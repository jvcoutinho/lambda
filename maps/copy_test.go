package maps_test

import (
	"fmt"
	"testing"

	"github.com/jvcoutinho/lambda/maps"
	"github.com/stretchr/testify/require"
)

func ExampleCopy() {
	age := map[string]int{
		"Bob":   25,
		"Alice": 19,
	}

	fmt.Println(maps.Copy(age))

	// Output:
	// map[Alice:19 Bob:25]
}

func TestCopy(t *testing.T) {
	t.Parallel()

	type TestCase[K comparable, V any] struct {
		description string
		currentMap  map[K]V
	}

	tests := []TestCase[string, int]{
		{
			description: "GivenMapIsNil_ShouldReturnNil",
			currentMap:  nil,
		},
		{
			description: "GivenMapIsNotNil_ShouldReturnCopy",
			currentMap:  map[string]int{"2": 2, "3": 3, "50": 50},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()

			// Act
			actualResult := maps.Copy(test.currentMap)

			// Assert
			require.Equal(t, test.currentMap, actualResult)
			require.NotSame(t, test.currentMap, actualResult)
		})
	}
}
