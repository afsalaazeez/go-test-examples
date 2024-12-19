package normal

import "testing"


func TestSum(t *testing.T) {

	type test struct {
		name     string
		x        int
		y        int
		expected int
	}

	tests := []test{
		{
			name:     "Normal operation with positive integers",
			x:        4,
			y:        5,
			expected: 9,
		},
		{
			name:     "Normal operation with negative integers",
			x:        -4,
			y:        -5,
			expected: -9,
		},
		{
			name:     "Normal operation with zero",
			x:        0,
			y:        5,
			expected: 5,
		},
		{
			name:     "Normal operation with large integers",
			x:        1000000,
			y:        2000000,
			expected: 3000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := Sum(tt.x, tt.y)

			if got != tt.expected {
				t.Errorf("Sum(%d, %d) = %d; want %d", tt.x, tt.y, got, tt.expected)
			}
		})
	}
}
