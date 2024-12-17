package parallel

import (
	"math"
	"testing"
)

func TestSum(t *testing.T) {

	tests := []struct {
		name    string
		x       int
		y       int
		want    int
		wantErr bool
	}{
		{
			name:    "Sum of Two Positive Integers",
			x:       5,
			y:       10,
			want:    15,
			wantErr: false,
		},
		{
			name:    "Sum of Two Negative Integers",
			x:       -5,
			y:       -3,
			want:    -8,
			wantErr: false,
		},
		{
			name:    "Sum of a Positive and a Negative Integer",
			x:       -5,
			y:       3,
			want:    -2,
			wantErr: false,
		},
		{
			name:    "Sum of Zero and an Integer",
			x:       0,
			y:       3,
			want:    3,
			wantErr: false,
		},
		{
			name:    "Sum of Maximum Integers",
			x:       math.MaxInt64,
			y:       math.MaxInt64,
			want:    -2,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := Sum(tt.x, tt.y)

			if (err != nil) != tt.wantErr {
				t.Errorf("Sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
