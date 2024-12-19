package normal

import "testing"

func TestSum(t *testing.T) {

	var testCases = []struct {
		x        int
		y        int
		expected int
	}{
		{1, 2, 3},
		{-1, -2, -3},
		{0, 5, 5},
		{1000000, 2000000, 3000000},
		{-5, 10, 5},
	}

	for _, tc := range testCases {

		result := Sum(tc.x, tc.y)

		if result != tc.expected {
			t.Errorf("Sum of (%v, %v): expected %v, got %v", tc.x, tc.y, tc.expected, result)
		} else {
			t.Logf("Sum of (%v, %v): expected %v, got %v", tc.x, tc.y, tc.expected, result)
		}
	}
}
