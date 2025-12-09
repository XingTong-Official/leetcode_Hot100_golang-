package main

import "fmt"

func main() {
	searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3)
}
func searchMatrix(matrix [][]int, target int) bool {
	lenY := len(matrix)
	// lenX := len(matrix[0])

	Y := 0
	up := 0
	down := lenY
	for {
		local := (down-up)/2 + up
		if up == down {
			Y = up
		}
		if matrix[local][0] == target {
			return true
		} else if matrix[local][0] > target {
			down = (down-up)/2 + up
		} else {
			up = (down-up)/2 + up + 1
		}
	}
	fmt.Println(Y)
	return true
}
