package normal

import (
	"testing"
)

// TestSum is a test function for the Sum function defined in the normal package.
// It uses table-driven tests to validate the Sum function against multiple scenarios.
/*func TestSum(t *testing.T) { // This TestSum function is redeclared in this block, so commenting it out.
	// Define test cases
	testCases := []struct {
		name        string
		x           int
		y           int
		expectedSum int
	}{
		{
			name:        "Scenario 1: One positive and one negative number",
			x:           5,
			y:           -3,
			expectedSum: 2,
		},
		{
			name:        "Scenario 2: Both numbers are zero",
			x:           0,
			y:           0,
			expectedSum: 0,
		},
		{
			name:        "Scenario 3: One number is zero and the other is a positive integer",
			x:           0,
			y:           7,
			expectedSum: 7,
		},
	}

	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function with test case parameters
			actualSum := Sum(tc.x, tc.y)

			// Check if the actual sum is equal to the expected sum
			if actualSum != tc.expectedSum {
				t.Errorf("Failed %s: Expected %d but got %d", tc.name, tc.expectedSum, actualSum)
			} else {
				t.Logf("Success %s: Expected %d and got %d", tc.name, tc.expectedSum, actualSum)
			}
		})
	}
}*/

// Sum is a function that adds two integers and returns the sum.
func Sum(x, y int) int {
	z := x + y
	return z
}
