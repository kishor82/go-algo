package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Already sorted slice",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted slice",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Mixed order slice",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
		},
		{
			name:     "Slice with duplicate elements",
			input:    []int{3, 3, 3, 2, 2, 1, 1},
			expected: []int{1, 1, 2, 2, 3, 3, 3},
		},
		{
			name:     "Slice with negative numbers",
			input:    []int{-3, 0, 5, -2, 1, -1, 4},
			expected: []int{-3, -2, -1, 0, 1, 4, 5},
		},
		{
			name:     "Single element slice",
			input:    []int{42},
			expected: []int{42},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.input); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.expected)
			}
		})
	}
}
