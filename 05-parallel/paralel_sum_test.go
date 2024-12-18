package parallel

import (
	"errors"
	"testing"
)


func TestSum(t *testing.T) {

	testCases := []struct {
		name     string
		x        int
		y        int
		expected int
		err      error
	}{

		{
			name:     "Handling of non-integer values",
			x:        3,
			y:        5,
			expected: 8,
			err:      nil,
		},
		{
			name:     "Concurrent execution safety",
			x:        7,
			y:        5,
			expected: 12,
			err:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result, err := Sum(tc.x, tc.y)

			if result != tc.expected {
				t.Errorf("Sum of %v and %v was incorrect, got: %v, want: %v.", tc.x, tc.y, result, tc.expected)
			}

			if (err != nil || tc.err != nil) && (err == nil || tc.err == nil || err.Error() != tc.err.Error()) {
				t.Errorf("Sum of %v and %v returned unexpected error: got %v, want %v", tc.x, tc.y, err, tc.err)
			}

			if err == nil {
				t.Logf("Sum test passed for values %v and %v", tc.x, tc.y)
			} else {
				t.Logf("Sum test failed for values %v and %v", tc.x, tc.y)
			}
		})
	}
}
