package main

import "fmt"

func main() {
	matrix := [][]int{{1, 1}}
	target := 5
	answer := searchMatrix(matrix, target)
	fmt.Println(answer)
}
func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	for {
		if j < 0 || i > len(matrix)-1 {
			return false
		}
		locate := matrix[i][j]
		if locate > target {
			j--
		} else if locate < target {
			i++
		} else {
			return true
		}
	}
}
