package happy_number

import "testing"

func TestHappyNumber(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{19, true},
		{2, false},
	}

	for _, test := range tests {
		if res := happyNumber(test.n); res != test.want {
			t.Errorf("happyNumber(%v) = %v, want %v", test.n, res, test.want)
		}
	}
}
