package main

import "fmt"

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
func largestRectangleArea(heights []int) int {
	stack := []int{-1}
	m := 0
	heights = append(heights, -1)
	for right, h := range heights {
		//循环出栈，保持单调栈
		for len(stack) > 1 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			m = max((right-left-1)*height, m)
		}
		stack = append(stack, right)
	}
	return m
}
