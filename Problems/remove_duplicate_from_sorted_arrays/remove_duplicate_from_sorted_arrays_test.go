package remove_duplicate_from_sorted_arrays

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 2},
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "All duplicates",
			nums:     []int{2, 2, 2, 2},
			expected: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeDuplicates(tt.nums)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}
