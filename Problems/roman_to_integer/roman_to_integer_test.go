package roman_to_integer

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXC", 1990},
		{"MCMXCIV", 1994},
		{"MMXXIV", 2024},
	}
	for _, test := range tests {
		result := romanToInt(test.input)
		if result != test.expected {
			t.Errorf("romanToInt(%s) = %d; expected %d", test.input, result, test.expected)
		}
	}
}
