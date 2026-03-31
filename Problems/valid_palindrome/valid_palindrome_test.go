package valid_palindrome

import "testing"

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}
	for _, test := range tests {
		result := validPalindrome(test.input)
		if result != test.expected {
			t.Errorf("validPalindrome(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}
