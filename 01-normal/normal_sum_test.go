package normal

import "testing"


func TestSum(t *testing.T) {

	testCases := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{
			name: "Normal operation with positive integers",
			x:    3,
			y:    5,
			want: 8,
		},
		{
			name: "Normal operation with negative integers",
			x:    -3,
			y:    -5,
			want: -8,
		},
		{
			name: "Normal operation with zero",
			x:    0,
			y:    5,
			want: 5,
		},
		{
			name: "Normal operation with large integers",
			x:    1000000,
			y:    2000000,
			want: 3000000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Sum(tc.x, tc.y)
			if got != tc.want {
				t.Errorf("Sum(%v, %v) = %v; want %v", tc.x, tc.y, got, tc.want)
			} else {
				t.Logf("Sum(%v, %v) = %v; passed", tc.x, tc.y, got)
			}
		})
	}
}
