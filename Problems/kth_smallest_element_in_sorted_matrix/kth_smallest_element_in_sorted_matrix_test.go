package kth_smallest_element_in_sorted_matrix

import "testing"

func TestKthSmallestElementInSortedMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		k      int
		want   int
	}{
		{[][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8, 13},
	}
	for _, tt := range tests {
		res := kthSmallestElementInSortedMatrix(tt.matrix, tt.k)
		if res != tt.want {
			t.Errorf("kthSmallestElementInSortedMatrix(%v,%v)=%v, got %v)", tt.matrix, tt.k, tt.want, res)
		}
	}
}
