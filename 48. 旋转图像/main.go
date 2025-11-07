package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix)
}
func rotate(matrix [][]int) {
	length := len(matrix)
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix)-i*2-1; j++ {
			//matrix[i][i+j],matrix[i+j][length-i-1],matrix[length-i-1][length-j-i-1],matrix[length-i-j-1][i]
			nowPoint := matrix[i][i+j]
			matrix[i][i+j] = matrix[length-i-j-1][i]
			matrix[length-i-j-1][i] = matrix[length-i-1][length-j-i-1]
			matrix[length-i-1][length-j-i-1] = matrix[i+j][length-i-1]
			matrix[i+j][length-i-1] = nowPoint
		}
	}
}
