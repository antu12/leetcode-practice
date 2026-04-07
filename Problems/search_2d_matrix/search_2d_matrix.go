package search_2d_matrix

func search2dMatrix(matrix [][]int, target int) bool {
	left := 0
	right := (len(matrix) * len(matrix[0])) - 1

	for left <= right {
		mid := (left + right) / 2
		row := mid / len(matrix[0])
		col := mid % len(matrix[0])

		res := matrix[row][col]
		if res == target {
			return true
		}
		if mid < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
