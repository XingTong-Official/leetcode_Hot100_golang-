package main

func main() {
	searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3)
}
func searchMatrix(matrix [][]int, target int) bool {
	x := len(matrix[0]) - 1
	y := 0
	for x >= 0 && y < len(matrix) {
		if matrix[y][x] == target {
			return true
		}
		if matrix[y][x] > target {
			x--
		} else {
			y++
		}
	}
	return false
}
