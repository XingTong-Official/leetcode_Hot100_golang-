package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}, {17, 18, 19, 20}, {21, 22, 23, 24}}
	answer := spiralOrder(matrix)
	fmt.Println(answer)
}
func spiralOrder(matrix [][]int) []int {
	length := len(matrix) * len(matrix[0])
	answer := []int{}
	i, j := 0, 0
	left, right := -1, len(matrix[0])
	up, down := -1, len(matrix)
	status := 1
	for {
		answer = append(answer, matrix[j][i])
		if len(answer) == length {
			break
		}
		if i+1 == right && j-1 == up && status == 1 {
			status = 2
			up++
		}
		if i+1 == right && j+1 == down && status == 2 {
			status = 3
			right--
		}
		if i-1 == left && j+1 == down && status == 3 {
			status = 4
			down--
		}
		if i-1 == left && j-1 == up && status == 4 {
			status = 1
			left++
		}
		switch status {
		case 1:
			i++
		case 2:
			j++
		case 3:
			i--
		case 4:
			j--
		}

	}
	return answer
}
