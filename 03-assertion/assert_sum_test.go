package assertion

import (
	"errors"
	"math"
	"testing"
)


func TestSum(t *testing.T) {

	testCases := []struct {
		name        string
		x           int
		y           int
		expectedSum int
		shouldError bool
	}{
		{
			name:        "Sum exceeds maximum integer limit",
			x:           math.MaxInt64,
			y:           1,
			expectedSum: 0,
			shouldError: true,
		},
		{
			name:        "Sum is below minimum integer limit",
			x:           math.MinInt64,
			y:           -1,
			expectedSum: 0,
			shouldError: true,
		},
		{
			name:        "Normal case",
			x:           5,
			y:           10,
			expectedSum: 15,
			shouldError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			sum, err := Sum(tc.x, tc.y)

			if (err != nil) != tc.shouldError {
				t.Errorf("Expected error: %v, but received error: %v", tc.shouldError, err)
			}

			if !tc.shouldError && sum != tc.expectedSum {
				t.Errorf("Expected sum: %v, but received: %v", tc.expectedSum, sum)
			}
		})
	}
}
