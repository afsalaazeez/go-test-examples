package normal

import (
	"math"
	"testing"
)

func TestSum(t *testing.T) {

	testCases := []struct {
		name        string
		x           int
		y           int
		expectedSum int
		expectError bool
	}{
		{
			"Normal operation with positive integers",
			5,
			7,
			12,
			false,
		},
		{
			"Normal operation with negative integers",
			-5,
			-7,
			-12,
			false,
		},
		{
			"Normal operation with zero",
			0,
			7,
			7,
			false,
		},
		{
			"Normal operation with large integers",
			1000000,
			2000000,
			3000000,
			false,
		},
		{
			"Error Handling",
			math.MaxInt64,
			1,
			0,
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result := Sum(tc.x, tc.y)

			if !tc.expectError {
				if result != tc.expectedSum {
					t.Errorf("Sum(%v, %v); got %v, want %v", tc.x, tc.y, result, tc.expectedSum)
				}
			} else {

				if result != tc.expectedSum {
					t.Errorf("Sum(%v, %v); expected an overflow error, but got %v", tc.x, tc.y, result)
				}
			}
		})
	}
}
