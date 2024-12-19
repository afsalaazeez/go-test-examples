package normal

import (
	"math"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/username/repo/normal"
)

func TestSum(t *testing.T) {

	testCases := []struct {
		name string
		x    int
		y    int
		sum  int
	}{
		{
			name: "Normal operation with positive integers",
			x:    5,
			y:    10,
			sum:  15,
		},
		{
			name: "Normal operation with negative integers",
			x:    -5,
			y:    -10,
			sum:  -15,
		},
		{
			name: "Normal operation with zero",
			x:    0,
			y:    10,
			sum:  10,
		},
		{
			name: "Normal operation with large integers",
			x:    math.MaxInt64 - 1,
			y:    1,
			sum:  math.MaxInt64,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := normal.Sum(tc.x, tc.y)

			assert.Equal(t, tc.sum, result, "they should be equal")
		})
	}

	t.Run("Error Handling", func(t *testing.T) {
		t.Log("Error Handling")

		result := normal.Sum(math.MaxInt64, 1)

		assert.True(t, result < math.MaxInt64, "result should be less than MaxInt64 due to overflow")
		assert.True(t, result < 1, "result should be less than 1 due to overflow")
	})
}
