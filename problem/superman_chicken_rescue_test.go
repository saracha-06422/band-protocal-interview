package problem

import (
	"testing"
)

func TestMaxChickensCovered(t *testing.T) {
	testCases := []struct {
		name              string
		numOfChickens     int
		distance          int
		chickensPositions []int
		expected          int
	}{
		{
			name:              "Example Case 1",
			numOfChickens:     5,
			distance:          5,
			chickensPositions: []int{2, 5, 10, 12, 15},
			expected:          2,
		},
		{
			name:              "Example Case 2",
			numOfChickens:     6,
			distance:          10,
			chickensPositions: []int{1, 11, 30, 34, 35, 37},
			expected:          4,
		},
		{
			name:              "Roof covering all chickens",
			numOfChickens:     5,
			distance:          20,
			chickensPositions: []int{1, 2, 3, 4, 5},
			expected:          5,
		},
		{
			name:              "Roof covering no chickens",
			numOfChickens:     3,
			distance:          0,
			chickensPositions: []int{10, 20, 30},
			expected:          1,
		},
		{
			name:              "Multiple chickens at the same position",
			numOfChickens:     6,
			distance:          1,
			chickensPositions: []int{5, 5, 5, 5, 5, 5},
			expected:          6,
		},
		{
			name:              "Large number of chickens, large distance",
			numOfChickens:     7,
			distance:          100000,
			chickensPositions: []int{1, 100000, 200000, 300000, 400000, 500000, 600000},
			expected:          2,
		},
		{
			name:              "Random positions",
			numOfChickens:     8,
			distance:          12,
			chickensPositions: []int{2, 8, 15, 18, 22, 25, 28, 35},
			expected:          4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := MaxChickensCovered(tc.chickensPositions, tc.numOfChickens, tc.distance)
			if result != tc.expected {
				t.Errorf("got : %d, expected : %d", result, tc.expected)
			}
		})
	}
}
