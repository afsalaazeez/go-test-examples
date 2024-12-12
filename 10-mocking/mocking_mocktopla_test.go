package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository is an autogenerated mock type for the MockRepository type
type MockRepository struct {
	mock.Mock
}

// MockTopla is a mock function for addition operation
func (m *MockRepository) MockTopla(sayilar []int) (int, error) {
	toplam := 0
	for i := range sayilar {
		toplam = toplam + sayilar[i]
	}
	return toplam, nil
}

// TestMockRepositoryMockTopla is a test function for MockTopla function
func TestMockRepositoryMockTopla(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		input    []int
		expected int
		err      error
	}{
		{
			name:     "Test scenario where one input number is positive and the other is negative",
			input:    []int{5, -3},
			expected: 2,
			err:      nil,
		},
		{
			name:     "Test scenario where both input numbers are zero",
			input:    []int{0, 0},
			expected: 0,
			err:      nil,
		},
		{
			name:     "Test scenario where one input number is zero and the other is a positive integer",
			input:    []int{0, 5},
			expected: 5,
			err:      nil,
		},
	}

	// Create new instance of MockRepository
	mockRepo := new(MockRepository)

	// Run table-driven tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := mockRepo.MockTopla(tc.input)
			// Log test case details
			t.Log("Executing test case where input is ", tc.input)

			// Assert operation errors
			if err != nil {
				t.Errorf("Expected no error, but got %v", err)
			}

			// Assert operation results
			assert.Equal(t, tc.expected, result, "Expected and actual results do not match")
		})
	}
}
