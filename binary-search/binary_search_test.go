package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		needle   int
		expected bool
	}{
		{
			name:     "Find existing element (middle)",
			input:    []int{1, 3, 5, 7, 9},
			needle:   5,
			expected: true,
		},
		{
			name:     "Find existing element (first)",
			input:    []int{1, 3, 5, 7, 9},
			needle:   1,
			expected: true,
		},
		{
			name:     "Find existing element (last)",
			input:    []int{1, 3, 5, 7, 9},
			needle:   9,
			expected: true,
		},
		{
			name:     "Element not in array (too small)",
			input:    []int{1, 3, 5, 7, 9},
			needle:   0,
			expected: false,
		},
		{
			name:     "Element not in array (too large)",
			input:    []int{1, 3, 5, 7, 9},
			needle:   10,
			expected: false,
		},
		{
			name:     "Element not in array (in between)",
			input:    []int{1, 3, 5, 7, 9},
			needle:   4,
			expected: false,
		},
		{
			name:     "Empty array",
			input:    []int{},
			needle:   5,
			expected: false,
		},
		{
			name:     "Single element array (found)",
			input:    []int{5},
			needle:   5,
			expected: true,
		},
		{
			name:     "Single element array (not found)",
			input:    []int{5},
			needle:   3,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := binarySearch(tt.input, tt.needle)
			if result != tt.expected {
				t.Errorf("binarySearch(%v, %d) = %v; want %v", tt.input, tt.needle, result, tt.expected)
			}
		})
	}
}
