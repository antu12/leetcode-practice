package search_2d_matrix

import "testing"

func TestSearch2dMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{[][]int{{1, 3, 5}, {7, 9, 11}, {13, 15, 17}}, 9, true},
		{[][]int{{1, 3, 5}, {7, 9, 11}, {13, 15, 17}}, 6, false},
		{[][]int{{1}}, 1, true},
	}

	for _, tt := range tests {
		result := search2dMatrix(tt.matrix, tt.target)
		if result != tt.want {
			t.Errorf("search2dMatrix(%v, %v) = %v; want %v", tt.matrix, tt.target, result, tt.want)
		}
	}
}
