package assertion

import (
	"errors"
	"math"
	"testing"
)


func TestSum(t *testing.T) {

	tests := []struct {
		name    string
		x       int
		y       int
		want    int
		wantErr error
	}{
		{
			name:    "Overflow Test",
			x:       math.MaxInt64,
			y:       1,
			want:    0,
			wantErr: errors.New("integer overflow/underflow"),
		},
		{
			name:    "Underflow Test",
			x:       math.MinInt64,
			y:       -1,
			want:    0,
			wantErr: errors.New("integer overflow/underflow"),
		},
		{
			name:    "Normal Test",
			x:       5,
			y:       10,
			want:    15,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sum(tt.x, tt.y)
			if got != tt.want || (err != nil && tt.wantErr == nil) || (err == nil && tt.wantErr != nil) {
				t.Errorf("Sum(%v, %v) = %v, %v; want %v, %v", tt.x, tt.y, got, err, tt.want, tt.wantErr)
			}
		})
	}
}
