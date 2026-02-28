package valid_parenthesis

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}
	for _, test := range tests {
		result := isValid(test.input)
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %t, Got: %t", test.input, test.expected, result)
		}
	}
}
