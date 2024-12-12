package table

import (
	"testing"
)

// TestSum is a test function for the Sum function defined in the same package.
// It uses table-driven testing to validate the Sum function against multiple test cases.
func TestSum(t *testing.T) {
	// Define a data structure for test cases
	type testCase struct {
		x        int
		y        int
		expected int
	}

	// Initialize the test cases
	testCases := []testCase{
		{x: 5, y: -2, expected: 3},   // Scenario 1: One positive number and one negative number
		{x: 0, y: 0, expected: 0},    // Scenario 2: Both numbers are zero
		{x: 0, y: 5, expected: 5},    // Scenario 3: One number is zero and the other is a positive integer
	}

	// Iterate over the test cases
	for i, tc := range testCases {
		// Run the function with the test case inputs and validate the output
		result := Sum(tc.x, tc.y)
		if result != tc.expected {
			t.Errorf("TestSum(%d): expected %d, got %d", i, tc.expected, result)
		} else {
			t.Logf("TestSum(%d): success", i)
		}
	}
}

func Sum(x, y int) int {
	z := x + y
	return z
}
