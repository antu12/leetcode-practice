package to_sum

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	result := twoSum(nums, target)
	if result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	nums = []int{3, 2, 4}
	target = 6
	expected = []int{1, 2}
	result = twoSum(nums, target)
	if result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	nums = []int{3, 3}
	target = 6
	expected = []int{0, 1}
	result = twoSum(nums, target)
	if result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
