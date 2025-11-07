package main

import "fmt"

func main() {
	nums := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(nums)
	fmt.Println(nums)
}
func setZeroes(matrix [][]int) {
	row, col := false, false
	for _, num := range matrix[0] {
		if num == 0 {
			row = true
			break
		}
	}
	for _, c := range matrix {
		if c[0] == 0 {
			col = true
			break
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < len(matrix[i]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}
	if row {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
	if col {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
