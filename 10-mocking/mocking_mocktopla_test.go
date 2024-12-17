package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func (i *islem) MockTopla(sayilar []int) (int, error) {
	toplam := 0
	for _, sayi := range sayilar {
		toplam += sayi
	}
	return toplam, nil
}
func TestMockRepositoryMockTopla(t *testing.T) {

	testCases := []struct {
		name     string
		input    []int
		expected int
		err      error
	}{
		{
			name:     "Regular Summation of Positive Integers",
			input:    []int{1, 2, 3, 4, 5},
			expected: 15,
			err:      nil,
		},
		{
			name:     "Summation Including Negative Integers",
			input:    []int{1, -2, 3, -4, 5},
			expected: 3,
			err:      nil,
		},
		{
			name:     "Summation of Zero",
			input:    []int{0, 0, 0, 0, 0},
			expected: 0,
			err:      nil,
		},
		{
			name:     "Summation of Empty Slice",
			input:    []int{},
			expected: 0,
			err:      nil,
		},
	}

	islemInstance := &islem{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result, err := islemInstance.MockTopla(tc.input)

			assert.Equal(t, tc.expected, result, "Expected output does not match the actual output")

			assert.Equal(t, tc.err, err, "Expected error does not match the actual error")
		})
	}
}
