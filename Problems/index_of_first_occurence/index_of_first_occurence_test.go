package index_of_first_occurence

import "testing"

func TestIndexOfFirstOccurence(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		expected int
	}{
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
	}
	for _, test := range tests {
		result := indexOfFirstOccurence(test.haystack, test.needle)
		if result != test.expected {
			t.Errorf("Haystack: %s, Needle: %s, Expected: %d, Got: %d", test.haystack, test.needle, test.expected, result)
		}
	}
}
