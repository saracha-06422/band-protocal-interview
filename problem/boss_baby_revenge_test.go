package problem

import (
	"testing"
)

func TestBossBabyRevenge2(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		name     string
	}{
		{"SRSSRRR", "Good boy", "Example Case 1"},
		{"RSSRR", "Bad boy", "Example Case 2"},
		{"SSSRRRRS", "Bad boy", "Example Case 3"},
		{"SRRSSR", "Bad boy", "Example Case 4"},
		{"SSRSRRR", "Good boy", "Example Case 5"},
		{"S", "Bad boy", "Single shot, no revenge"},
		{"SR", "Good boy", "Immediate revenge"},
		{"R", "Bad boy", "Invalid start with revenge"},
		{"SSRR", "Good boy", "Balanced shots and revenge"},
		{"SRRR", "Good boy", "Extra revenge actions"},
		{"SSSRRR", "Good boy", "All shots revenged"},
		{"SRSSR", "Bad boy", "Unrevenged shot at the end"},
		{"SSSSRRRR", "Good boy", "Late revenge but all shots revenged"},
		{"RRRR", "Bad boy", "No shots, only revenge"},
		{"SSSS", "Bad boy", "Only shots, no revenge"},
		{"SRSRSR", "Good boy", "Alternating shots and revenge"},
		{"RSSS", "Bad boy", "Invalid initial revenge"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := BossBabyRevenge(tc.input)
			if result != tc.expected {
				t.Errorf("For input '%s', expected '%s' but got '%s'", tc.input, tc.expected, result)
			}
		})
	}
}
