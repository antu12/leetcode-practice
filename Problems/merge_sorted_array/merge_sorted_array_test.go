package merge_sorted_array

import "testing"

func TestMergeSortedArray(t *testing.T) {
	tests := []struct {
		name string
		num1 []int
		m    int
		num2 []int
		n    int
		want []int
	}{
		{
			name: "test1",
			num1: []int{1, 2, 3, 0, 0, 0},
			m:    3,
			num2: []int{2, 5, 6},
			n:    3,
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "test2",
			num1: []int{1},
			m:    1,
			num2: []int{},
			n:    0,
			want: []int{1},
		},
		{
			name: "test3",
			num1: []int{0},
			m:    0,
			num2: []int{1},
			n:    1,
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mergeSortedArray(tt.num1, tt.m, tt.num2, tt.n)
			for i := range tt.want {
				if tt.num1[i] != tt.want[i] {
					t.Errorf("mergeSortedArray() = %v, want %v", tt.num1, tt.want)
					break
				}
			}
		})
	}
}
