package assertion

import (
	"testing"
	"errors"
)


func TestSum(t *testing.T) {

	tests := []struct {
		name     string
		x        int
		y        int
		expected int
	}{
		{
			name:     "Test with positive integers",
			x:        5,
			y:        3,
			expected: 8,
		},
		{
			name:     "Test with zero",
			x:        5,
			y:        0,
			expected: 5,
		},
		{
			name:     "Test with negative integers",
			x:        -5,
			y:        -3,
			expected: -8,
		},
		{
			name:     "Test with mixed positive and negative integers",
			x:        5,
			y:        -3,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Sum(tt.x, tt.y)
			if err != nil {
				t.Errorf("Sum() error = %v, expected error = nil", err)
				return
			}
			if result != tt.expected {
				t.Errorf("Sum() = %v, expected = %v", result, tt.expected)
			}
		})
	}

	t.Run("Test error handling", func(t *testing.T) {
		_, err := Sum(1, 1)
		if err != nil {
			t.Errorf("Sum() error = %v, expected error = nil", err)
		}
	})
}
