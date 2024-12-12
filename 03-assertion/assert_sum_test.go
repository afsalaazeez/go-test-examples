package assertion

import (
	"testing"
)

// Function for which test cases are written
func Sum(x, y int) (int, error) {
	z := x + y
	return z, nil
}

// TestSum is a test function for Sum
func TestSum(t *testing.T) {
	// Define test cases
	testCases := []struct {
		desc     string
		in1      int
		in2      int
		expected int
	}{
		{
			desc:     "Test scenario where one input number is positive and the other is negative",
			in1:      5,
			in2:      -3,
			expected: 2,
		},
		{
			desc:     "Test scenario where both input numbers are zero",
			in1:      0,
			in2:      0,
			expected: 0,
		},
		{
			desc:     "Test scenario where one input number is zero and the other is a positive integer",
			in1:      0,
			in2:      5,
			expected: 5,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := Sum(tc.in1, tc.in2)
			if err != nil {
				t.Fatalf("Sum() returned an error: %v", err)
			}
			if result != tc.expected {
				t.Errorf("Sum() = %v; want %v", result, tc.expected)
			} else {
				t.Logf("Success: %s", tc.desc)
			}
		})
	}
}
