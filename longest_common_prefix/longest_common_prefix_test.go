package longest_common_prefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	expected := "fl"
	result := longestCommonPrefix(strs)
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	strs = []string{"dog", "racecar", "car"}
	expected = ""
	result = longestCommonPrefix(strs)
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	strs = []string{"interspecies", "interstellar", "interstate"}
	expected = "inters"
	result = longestCommonPrefix(strs)
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
