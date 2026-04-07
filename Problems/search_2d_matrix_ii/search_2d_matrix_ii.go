package search_2d_metrix_ii

func search2dMetrixIi(matrix [][]int, target int) bool {
	row := 0
	col := len(matrix[0]) - 1

	for row < len(matrix) && col >= 0 {
		res := matrix[row][col]

		if res == target {
			return true
		}

		if res > target {
			col--
		} else {
			row++
		}
	}
	return false
}
