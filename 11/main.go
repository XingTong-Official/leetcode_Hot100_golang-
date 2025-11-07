package main

import "fmt"

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	answer := maxArea(nums)
	fmt.Println(answer)
}
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	m := 0
	for left != right {
		heig := min(height[left], height[right])
		width := right - left
		m = max(heig*width, m)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return m
}
