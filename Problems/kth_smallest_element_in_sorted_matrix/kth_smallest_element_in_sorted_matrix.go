package kth_smallest_element_in_sorted_matrix

func kthSmallestElementInSortedMatrix(matrix [][]int, k int) int {
	n := len(matrix)

	l := matrix[0][0]
	r := matrix[n-1][n-1]

	for l < r {
		mid := (l + r) / 2

		count := 0
		col := n - 1

		for row := 0; row < n; row++ {
			if col >= 0 && matrix[row][col] > mid {
				col--
			}
			count += col + 1
		}

		if count < k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
