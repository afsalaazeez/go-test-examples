package parallel

import (
	"testing"
)

// Unit test for Sum function
func TestSum(t *testing.T) {
	// Table driven tests
	var tests = []struct {
		name  string
		x     int
		y     int
		want  int
	}{
		{"Positive and Negative", 5, -3, 2},
		{"Both Zeros", 0, 0, 0},
		{"Zero and Positive", 0, 4, 4},
	}
	// Loop over test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sum(tt.x, tt.y)
			// If error is reported while there should not be
			if err != nil {
				t.Fatalf("Sum(%v, %v) returned unexpected error: %v", tt.x, tt.y, err)
			}
			// If the result does not match the expected
			if got != tt.want {
				t.Fatalf("Sum(%v, %v) = %v, want %v", tt.x, tt.y, got, tt.want)
			}
			t.Logf("Scenario %v passed", tt.name)
		})
	}
}
