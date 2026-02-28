package remove_element

import "testing"

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums     []int
		val      int
		expected int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}
	for _, test := range tests {
		result := removeElement(test.nums, test.val)
		if result != test.expected {
			t.Errorf("removeElement(%v, %d) = %d; expected %d", test.nums, test.val, result, test.expected)
		}
	}
}
