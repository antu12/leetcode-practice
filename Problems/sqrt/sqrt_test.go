package sqrt

import "testing"

func TestSqrt(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{4, 2},
		{8, 2},
		{0, 0},
		{1, 1},
		{15, 3},
		{16, 4},
		{17, 4},
	}
	for _, tt := range tests {
		if got := sqrt(tt.x); got != tt.want {
			t.Errorf("sqrt(%d) = %d; want %d", tt.x, got, tt.want)
		}
	}
}
