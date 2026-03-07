package add_binary

import "testing"

func TestAddBinary(t *testing.T) {
	tests := []struct {
		a, b   string
		result string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}

	for _, test := range tests {
		if res := addBinary(test.a, test.b); res != test.result {
			t.Errorf("addBinary(%s, %s) = %s; want %s", test.a, test.b, res, test.result)
		}
	}
}
