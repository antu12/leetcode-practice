package search_insert_position

import "testing"

func TestSearchInsertPosition(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{nums: []int{1, 3, 5, 6}, target: 5, expected: 2},
		{nums: []int{1, 3, 5, 6}, target: 2, expected: 1},
		{nums: []int{1, 3, 5, 6}, target: 7, expected: 4},
	}

	for _, test := range tests {
		result := searchInsertPosition(test.nums, test.target)
		if result != test.expected {
			t.Errorf("For input nums: %v and target: %d, expected %d but got %d", test.nums, test.target, test.expected, result)
		}
	}
}
