package search_2d_metrix_ii

import "testing"

func TestSearch2dMetrixIi(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{[][]int{{1, 4, 7, 11}, {2, 5, 8, 12}, {3, 6, 9, 16}, {10, 13, 14, 17}}, 9, true},
	}
	for _, tt := range tests {
		res := search2dMetrixIi(tt.matrix, tt.target)
		if res != tt.want {
			t.Errorf("search2dMatrixIi(%v, %v) = %v; want %v", tt.matrix, tt.target, res, tt.want)
		}
	}
}
